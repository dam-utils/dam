package run

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"dam/config"
	"dam/driver/db"
	"dam/driver/decorate"
	"dam/driver/engine"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/driver/structures"
	"dam/run/internal"
)

func InstallApp(appCurrentName string) {
	var isInstallingFromFile bool

	if fs.IsExistFile(appCurrentName) {
		isInstallingFromFile = true
		flag.ValidateFilePath(appCurrentName)
	} else {
		isInstallingFromFile = false
		flag.ValidateAppPlusVersion(appCurrentName)
	}
	logger.Debug("Flags validated with success")
	logger.Success("Start '%s' installing to the system.", appCurrentName)

	logger.Debug("Preparing docker image ...")
	var tag string
	if isInstallingFromFile {
		tag = getTagFromArchiveManifest(appCurrentName)
		engine.VDriver.LoadImage(appCurrentName)
	} else {
		tag = dockerPull(appCurrentName)
	}

	logger.Debug("Preparing family label ...")
	familyLabel := internal.GetFamily(tag)

	logger.Debug("Preparing multiversion label ...")
	if !internal.GetMultiVersion(tag) {
		logger.Warn("Not set multiversion flag for this app")
		isExistFamily(familyLabel)
	}

	logger.Debug("Getting meta ...")
	tmpDir := internal.PrepareTmpMetaPath(config.TMP_META_PATH)
	defer fs.Remove(tmpDir)

	containerId := engine.VDriver.ContainerCreate(tag, "")
	engine.VDriver.CopyFromContainer(containerId, string(os.PathSeparator)+config.META_DIR_NAME, tmpDir)
	engine.VDriver.ContainerRemove(containerId)

	logger.Debug("Printing description ...")
	decorate.PrintDescription(filepath.Join(tmpDir, config.META_DIR_NAME, config.DESCRIPTION_FILE_NAME))

	logger.Debug("Installing image ...")
	installMeta := filepath.Join(tmpDir, config.META_DIR_NAME)
	install := getInstall(installMeta)
	fs.RunFile(install)

	logger.Debug("Saving to DB ...")
	saveAppToDB(tag, familyLabel)
	logger.Success("App '%s' was installed.", appCurrentName)
}

func isExistFamily(imageFamily string) {
	if db.ADriver.ExistFamily(imageFamily) {
		logger.Fatal("Cannot add the application to DB. App with FAMILY '%s' is exist in DB", imageFamily )
	}
}

func dockerPull(app string) string {
	defRepo := db.RDriver.GetDefaultRepo()
	if defRepo == nil {
		logger.Fatal("Internal error. Not found default repo")
	}

	var tag string
	if defRepo.Id == structures.OfficialRepo.Id {
		tag = app
	} else {
		tag = defRepo.Server + "/" + app
	}

	if defRepo.Id == structures.OfficialRepo.Id {
		tag = app
	}

	engine.VDriver.Pull(tag, defRepo)

	return tag
}

func saveAppToDB(tag, familyLabel string) {
	newRepo, imageName, imageVersion := internal.SplitTag(tag)

	newId := structures.OfficialRepo.Id
	if newRepo == "" {
		newId = internal.PrepareRepo(newRepo)
	}

	var app structures.App
	app.RepoID = newId
	app.DockerID = engine.VDriver.GetImageID(tag)
	app.ImageName = imageName
	app.ImageVersion = imageVersion
	app.Family = familyLabel

	db.ADriver.NewApp(&app)
}

func getInstall(meta string) string {
	inst := filepath.Join(meta, config.INSTALL_FILE_NAME)
	if !fs.IsExistFile(inst) {
		logger.Fatal("Not found '%s' file in meta '%s'", config.INSTALL_FILE_NAME, config.META_DIR_NAME)
	}
	return inst
}

func getTagFromArchiveManifest(appCurrentName string) string {
	//TODO read manifest without archive uncompressing
	gzipFile := fs.Gunzip(appCurrentName)
	defer fs.Remove(gzipFile)
	tarGzipDir := fs.Untar(gzipFile)
	defer fs.Remove(tarGzipDir)

	manifestFile := tarGzipDir + string(filepath.Separator) + config.SAVE_MANIFEST_FILE

	content, err := os.Open(manifestFile)
	defer func() {
		if content != nil {
			content.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open the manifest file '%s' with error: %s", manifestFile, err)
	}

	type manifest struct {
		RepoTags []string `json:"RepoTags"`
	}

	result := make([]manifest, 0)
	byteValue, err := ioutil.ReadAll(content)
	if err != nil {
		logger.Fatal("Cannot read content in manifest file with error: %s", err)
	}

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		logger.Fatal("Cannot unmarshal manifest file with error: %s", err)
	}

	if len(result) > 0 {
		if len(result[0].RepoTags) > 0 {
			flag.ValidateAppPlusVersion(result[0].RepoTags[0])
			return result[0].RepoTags[0]
		}
	}

	logger.Fatal("Cannot get manifest tag from archive '%s'", appCurrentName)
	return ""
}

package run

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"dam/driver/conf/option"
	"dam/driver/db"
	"dam/driver/decorate"
	"dam/driver/engine"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/driver/structures"
	"dam/run/internal"
	"dam/run/internal/label/servers"
)

type installArgType int
const (
	unknownInstall installArgType = iota
	fileInstall
	appInstall
	tagInstall
)

func InstallApp(arg string) {
	installType := getInstallTypeByArg(arg)

	logger.Debug("Validating docker image with type '%v' ...", installType)
	validateInstallArg(arg, installType)
	logger.Debug("Flags validated with success")
	logger.Success("Start '%s' installing to the system.", arg)

	logger.Debug("Preparing docker image ...")
	var imageTag string
	switch installType {
	case fileInstall:
		imageTag = getTagFromArchiveManifest(arg)
		engine.VDriver.LoadImage(arg)
	case appInstall:
		imageTag = dockerPull(arg)
	case tagInstall:
		repo, _, _ := internal.SplitTag(arg)
		internal.PrepareRepo(repo)
		imageTag = arg
	default:
		logger.Fatal("Internal error. Not supported type of the install argument")
	}

	logger.Debug("Validate existing the image in the docker cache...")
	imageId := engine.VDriver.GetImageID(imageTag)
	if imageId == "" {
		logger.Fatal("Stop installing image with tag '%s'. Cannot find image in the docker cache.", imageTag)
	}

	familyLabel := prepareLabelsWithFamily(imageTag)

	logger.Debug("Getting meta ...")
	tmpDir := internal.PrepareTmpMetaPath(option.Config.FileSystem.GetTmpMetaPath())
	defer fs.Remove(tmpDir)

	containerId := engine.VDriver.ContainerCreate(imageTag, "")
	engine.VDriver.CopyFromContainer(containerId, string(os.PathSeparator)+option.Config.FileSystem.GetMetaDirName(), tmpDir)
	engine.VDriver.ContainerRemove(containerId)

	logger.Debug("Printing description ...")
	decorate.PrintDescription(filepath.Join(tmpDir, option.Config.FileSystem.GetMetaDirName(), option.Config.FileSystem.GetDescriptionFileName()))

	logger.Debug("Installing image ...")
	installMeta := filepath.Join(tmpDir, option.Config.FileSystem.GetMetaDirName())
	install := getInstall(installMeta)
	fs.RunFile(install)

	logger.Debug("Saving to DB ...")
	saveAppToDB(imageTag, familyLabel)
	logger.Success("App '%s' was installed.", imageTag)
}

func getInstallTypeByArg(arg string) installArgType {
	if fs.IsExistFile(arg) {
		return fileInstall
	}

	repo, name, version := internal.SplitTag(arg)
	logger.Debug("Split tag '%s' as Repo='%s', Name='%s' and Version='%s'", arg, repo, name, version)
	if repo != "" && name != "" && version != "" {
		return tagInstall
	}

	if name != "" && version != "" {
		return appInstall
	}

	return unknownInstall
}

func validateInstallArg(arg string, installType installArgType) {
	switch installType {
	case fileInstall:
		flag.ValidateFilePath(arg)
	case appInstall:
		flag.ValidateAppPlusVersion(arg)
	case tagInstall:
		flag.ValidateTag(arg)
	default:
		logger.Fatal("Unknown argument '%s' for command 'install'. See '%s help install'", arg, option.Config.Global.GetProjectName())
	}
}

func prepareLabelsWithFamily(imageTag string) string {
	logger.Debug("Preparing servers label ...")
	serversLabel := internal.GetServersByTag(imageTag)
	logger.Debug("Servers labels '%v' for tag '%s'", serversLabel, imageTag)
	createTagImages(imageTag, serversLabel)

	logger.Debug("Preparing family label ...")
	familyLabel := internal.GetFamily(imageTag)

	logger.Debug("Preparing multiversion label ...")
	if !internal.GetMultiVersion(imageTag) {
		logger.Warn("Not set multiversion flag for this app")
		isExistFamily(familyLabel)
	}
	return familyLabel
}

// Create tags different from the given one
func createTagImages(tag, serversLabel string) {
	defRepo, name, version := internal.SplitTag(tag)
	imageId := engine.VDriver.GetImageID(tag)
	if imageId == "" {
		logger.Fatal("Image with tag '%s' not exist in the system", tag)
	}

	storage := servers.NewLabel(serversLabel)
	err := storage.ValidateRepos()
	if err != nil {
		logger.Fatal("Failed validating servers label '%s' with error: %s", storage.String(), err)
	}
	storage.AddRepo(defRepo)

	reposList, official := storage.ReposList()
	for _, repo := range reposList {
		if repo != defRepo {
			prepareImageTag(imageId, repo+"/"+name+":"+version)
		}
	}

	if official && defRepo != "" {
		prepareImageTag(imageId, name+":"+version)
	}
}

func prepareImageTag(imageId, tag string) {
	newId := engine.VDriver.GetImageID(tag)
	if newId == "" {
		logger.Debug("Create tag '%s' for image id '%s'", tag, imageId)
		engine.VDriver.CreateTag(imageId, tag)
	}
}

func isExistFamily(imageFamily string) {
	if db.ADriver.ExistFamily(imageFamily) {
		logger.Fatal("Cannot add the application to DB. App with FAMILY '%s' was installed in the system.", imageFamily)
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

	newRepoId := structures.OfficialRepo.Id
	if newRepo != "" {
		newRepoId = internal.PrepareRepo(newRepo)
	}

	var app structures.App
	app.RepoID = newRepoId
	app.DockerID = engine.VDriver.GetImageID(tag)
	app.ImageName = imageName
	app.ImageVersion = imageVersion
	app.Family = familyLabel

	db.ADriver.NewApp(&app)
}

func getInstall(meta string) string {
	inst := filepath.Join(meta, option.Config.FileSystem.GetInstallFileName())
	if !fs.IsExistFile(inst) {
		logger.Fatal("Not found '%s' file in meta '%s'", option.Config.FileSystem.GetInstallFileName(), option.Config.FileSystem.GetMetaDirName())
	}
	return inst
}

func getTagFromArchiveManifest(appCurrentName string) string {
	//TODO read manifest without archive uncompressing
	gzipFile := fs.Gunzip(appCurrentName)
	defer fs.Remove(gzipFile)
	tarGzipDir := fs.Untar(gzipFile)
	defer fs.Remove(tarGzipDir)

	manifestFile := path.Join(tarGzipDir, option.Config.Save.GetManifestFile())

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

	logger.Debug("Parsed manifest file '%v'", result)
	if len(result) > 0 {
		if len(result[0].RepoTags) > 0 {
			return result[0].RepoTags[0]
		}
	}

	logger.Fatal("Cannot get manifest tag from archive '%s'", appCurrentName)
	return ""
}

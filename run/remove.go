package run

import (
	"dam/driver/engine"
	"dam/driver/structures"
	"os"
	"path/filepath"
	"strings"

	"dam/config"
	"dam/driver/db"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/run/internal"
)

type RemoveAppSettings struct {
	Force bool
}

var RemoveAppFlags = new(RemoveAppSettings)

func RemoveApp(name string) {
	flag.ValidateAppName(name)
	logger.Debug("Flags validated with success")

	logger.Debug("Getting app tag ...")
	app := getAppIdByName(name)
	tag := getTagFormApp(app)

	logger.Success("Start app '%s:%s' removing from the system.", app.ImageName, app.ImageVersion)

	logger.Debug("Getting meta ...")
	tmpMeta := internal.PrepareTmpMetaPath(config.TMP_META_PATH)
	defer fs.Remove(tmpMeta)

	logger.Debug("tmpMeta: '%v'", tmpMeta)
	containerId := engine.VDriver.ContainerCreate(tag, "")
	engine.VDriver.CopyFromContainer(containerId, string(os.PathSeparator)+config.META_DIR_NAME, tmpMeta)
	engine.VDriver.ContainerRemove(containerId)

	logger.Debug("Uninstalling image ...")
	uninstallMeta := filepath.Join(tmpMeta, config.META_DIR_NAME)
	uninstall := getUninstall(uninstallMeta)

	logger.Debug("Running uninstall ...")
	fs.RunFile(uninstall)

	logger.Debug("Removing app from DB ...")
	db.ADriver.RemoveAppById(app.Id)

	logger.Success("App '%s:%s' was removed.", app.ImageName, app.ImageVersion)
}

func getAppIdByName(name string) *structures.App {
	id := -1

	apps := db.ADriver.GetApps()
	for _, app := range apps {
		logger.Info("app.Name:'%s', name:'%s'", app.ImageName, name)
		if app.ImageName == name {
			id = app.Id
		}
	}

	if id == -1 {
		logger.Fatal("Not found app with name '%s' in DB", name)
	}

	app := db.ADriver.GetAppById(id)
	logger.Debug("Remove app '%s'", app)
	if app == nil {
		logger.Fatal("Not found app with name '%s' in DB", name)
	}

	return app
}

func getTagFormApp(app *structures.App) string {
	var tag strings.Builder

	if app.RepoID != structures.OfficialRepo.Id {
		repo := db.RDriver.GetRepoById(app.RepoID)
		if repo == nil {
			logger.Fatal("Internal error. Cannot get repo for ID '%v'", app.RepoID)
		}
		tag.WriteString(repo.Server)
		tag.WriteString("/")
	}
	tag.WriteString(app.ImageName)
	tag.WriteString(":")
	tag.WriteString(app.ImageVersion)

	return tag.String()
}

func getUninstall(meta string) string {
	uninst := filepath.Join(meta, config.UNINSTALL_FILE_NAME)
	if !fs.IsExistFile(uninst) {
		logger.Fatal("Not found '%s' file in meta '%s'", config.UNINSTALL_FILE_NAME, config.META_DIR_NAME)
	}
	return uninst
}

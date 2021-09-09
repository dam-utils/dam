package run

import (
	"os"
	"path/filepath"

	"dam/driver/conf/option"
	"dam/driver/db"
	"dam/driver/engine"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/structures"
	"dam/driver/validate"
	"dam/run/internal"
)

type RemoveAppSettings struct {
	Force bool
}

var RemoveAppFlags = new(RemoveAppSettings)

func RemoveApp(arg string) {
	var app *structures.App

	if validate.CheckApp(arg) == nil {
		app = getAppIdByTag(arg)
	} else {
		if validate.CheckAppName(arg) == nil {
			app = getAppIdByName(arg)
		} else {
			logger.Fatal("Cannot parse the command argument. It is not a tag or an app name.")
		}
	}
	logger.Debug("Flags validated with success")

	logger.Success("Start app '%s:%s' removing from the system.", app.ImageName, app.ImageVersion)
	defer func() {
		if RemoveAppFlags.Force {
			logger.Debug("Forced app removing from DB ...")
			db.ADriver.RemoveAppById(app.Id)
			logger.Success("App '%s:%s' was removed.", app.ImageName, app.ImageVersion)
		}
	}()

	logger.Debug("Getting meta ...")
	tmpMeta := internal.PrepareTmpMetaPath(option.Config.FileSystem.GetTmpMetaPath())
	defer fs.Remove(tmpMeta)

	logger.Debug("tmpMeta: '%v'", tmpMeta)
	containerId := engine.VDriver.ContainerCreate(app.DockerID, "")
	engine.VDriver.CopyFromContainer(containerId, string(os.PathSeparator)+option.Config.FileSystem.GetMetaDirName(), tmpMeta)
	engine.VDriver.ContainerRemove(containerId)

	logger.Debug("Uninstalling image ...")
	uninstallMeta := filepath.Join(tmpMeta, option.Config.FileSystem.GetMetaDirName())
	uninstall := getUninstall(uninstallMeta)

	logger.Debug("Running uninstall ...")
	fs.RunFile(uninstall)

	if !RemoveAppFlags.Force {
		logger.Debug("App removing from DB ...")
		db.ADriver.RemoveAppById(app.Id)

		logger.Success("App '%s:%s' was removed.", app.ImageName, app.ImageVersion)
	}
}

func getAppIdByName(name string) *structures.App {
	ids := make([]int,0)

	apps := db.ADriver.GetApps()
	for _, app := range apps {
		if app.ImageName == name {
			ids = append(ids, app.Id)
		}
	}

	var resultID int

	switch len(ids) {
	case 0:
		logger.Fatal("Not found app with name '%s' in DB", name)
	case 1:
		resultID = ids[0]
	default:
		logger.Fatal("Found many apps with name '%s' in DB. You need remove by <name>:<version>", name)
	}

	app := db.ADriver.GetAppById(resultID)
	logger.Debug("Remove app '%s'", app)
	if app == nil {
		logger.Fatal("Not found app with name '%s' and id '%v' in DB", name, resultID)
	}

	return app
}

func getAppIdByTag(tag string) *structures.App {
	id := -1
	_, name, version := internal.SplitTag(tag)

	apps := db.ADriver.GetApps()
	for _, app := range apps {
		if app.ImageName == name && app.ImageVersion == version {
			// Не делаю проверку, вдруг с таким тэгом приложений в базе несколько
			id = app.Id
		}
	}

	if id == -1 {
		logger.Fatal("Not found app with tag '%s'", tag)
	}

	app := db.ADriver.GetAppById(id)
	logger.Debug("Remove app '%s'", app)
	if app == nil {
		logger.Fatal("Not found app with tag '%s' and id '%v' in DB", tag, id)
	}

	return app
}

func getUninstall(meta string) string {
	uninstall := filepath.Join(meta, option.Config.FileSystem.GetUninstallFileName())
	if !fs.IsExistFile(uninstall) {
		logger.Fatal("Not found '%s' file in meta '%s'", option.Config.FileSystem.GetUninstallFileName(), option.Config.FileSystem.GetMetaDirName())
	}
	return uninstall
}

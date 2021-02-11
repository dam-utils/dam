package run

import (
	"os"
	"path/filepath"
	"strings"

	"dam/config"
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
	var tag string

	if validate.CheckApp(arg) == nil {
		app = getAppIdByTag(arg)
		tag = arg
	} else {
		if validate.CheckAppName(arg) == nil {
			app = getAppIdByName(arg)
			tag = getTagFormApp(app)
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

package internal

import (
	"os"
	"strconv"
	"strings"

	"dam/driver/conf/option"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/structures"
	"dam/driver/validate"
)

func SaveApps(apps []*structures.App) {
	f, err := os.OpenFile(option.Config.FilesDB.GetTmp(), os.O_WRONLY|os.O_CREATE, option.Config.FilesDB.GetFilesPermissions())
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open apps file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}

	for _, app := range apps {
		newLine := app2str(app)
		_, err := f.WriteString(*newLine)
		if err != nil {
			logger.Fatal("Cannot write to apps file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
		}
	}
	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync apps file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from apps file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}

	logger.Debug("Move '%s' to '%s'", option.Config.FilesDB.GetTmp(), option.Config.FilesDB.GetAppsFilename())
	fs.MoveFile(option.Config.FilesDB.GetTmp(), option.Config.FilesDB.GetAppsFilename())
}

func app2str(app *structures.App) *string {
	var appStr string
	sep := option.Config.FilesDB.GetSeparator()

	multiVers := ""
	if app.MultiVersion {
		multiVers = option.Config.Decoration.GetBoolFlagSymbol()
	}

	fields := []string{
		strconv.Itoa(app.Id),
		app.DockerID,
		app.ImageName,
		app.ImageVersion,
		strconv.Itoa(app.RepoID),
		multiVers,
		app.Family,
	}
	lenF := len(fields)
	for i, field := range fields {
		if i == lenF - 1 {
			appStr = appStr + field + "\n"
		} else {
			appStr = appStr + field + sep
		}
	}
	return &appStr
}

func GetNewAppID(apps []*structures.App) int {
	res := 0

	if len(apps) == 0 {
		return 0
	}
	for _, app := range apps {
		if app.Id >= res {
			res = app.Id
		}
	}
	return res + 1
}

func Str2app(str string) *structures.App {
	app := new(structures.App)
	strArray := strings.Split(str, option.Config.FilesDB.GetSeparator())

	if validate.CheckAppID(strArray[0]) != nil {
		logger.Fatal("Internal error. Cannot parse the app ID in line '%s'", str)
	}
	if validate.CheckDockerID(strArray[1]) != nil {
		logger.Fatal("Internal error. Cannot parse the docker ID in line '%s'", str)
	}
	if validate.CheckAppName(strArray[2]) != nil {
		logger.Fatal("Internal error. Cannot parse the app name in line '%s'", str)
	}
	if validate.CheckVersion(strArray[3]) != nil {
		logger.Fatal("Internal error. Cannot parse the app version in line '%s'", str)
	}
	if validate.CheckRepoID(strArray[4]) != nil {
		logger.Fatal("Internal error. Cannot parse the repo id in line '%s'", str)
	}
	if validate.CheckBool(strArray[5]) != nil {
		logger.Fatal("Internal error. Cannot parse the multiversion flag in line '%s'", str)
	}
	if validate.CheckLabel(strArray[6]) != nil {
		logger.Fatal("Internal error. Cannot parse the family flag in line '%s'", str)
	}

	app.Id, _ = strconv.Atoi(strArray[0])
	app.DockerID = strArray[1]
	app.ImageName = strArray[2]
	app.ImageVersion = strArray[3]
	app.RepoID, _ = strconv.Atoi(strArray[4])
	if strArray[5] == option.Config.FilesDB.GetBoolFlagSymbol() {
		app.MultiVersion = true
	}
	app.Family = strArray[6]
	return app
}

func ClearApps() {
	f, err := os.OpenFile(option.Config.FilesDB.GetTmp(), os.O_WRONLY|os.O_CREATE, option.Config.FilesDB.GetFilesPermissions())
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open apps file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}

	_, err = f.WriteString("")
	if err != nil {
		logger.Fatal("Cannot write to apps file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}

	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync apps file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from apps file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}

	logger.Debug("Move '%s' to '%s'", option.Config.FilesDB.GetTmp(), option.Config.FilesDB.GetAppsFilename())
	fs.MoveFile(option.Config.FilesDB.GetTmp(), option.Config.FilesDB.GetAppsFilename())
}
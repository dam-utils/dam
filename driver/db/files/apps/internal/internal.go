package internal

import (
	"os"
	"strconv"
	"strings"

	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/structures"
	"dam/driver/validate"
)

func SaveApps(apps []*structures.App) {
	f, err := os.OpenFile(config.FILES_DB_TMP_DIR, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open apps file '%s' with error: %s", config.FILES_DB_TMP_DIR, err)
	}

	for _, app := range apps {
		newLine := app2str(app)
		_, err := f.WriteString(*newLine)
		if err != nil {
			logger.Fatal("Cannot write to apps file '%s' with error: %s", config.FILES_DB_TMP_DIR, err)
		}
	}
	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync apps file '%s' with error: %s", config.FILES_DB_TMP_DIR, err)
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from apps file '%s' with error: %s", config.FILES_DB_TMP_DIR, err)
	}

	logger.Debug("Move '%s' to '%s'", config.FILES_DB_TMP_DIR, config.FILES_DB_APPS_FILENAME)
	fs.MoveFile(config.FILES_DB_TMP_DIR, config.FILES_DB_APPS_FILENAME)
}

func app2str(app *structures.App) *string {
	var appStr string
	sep := config.FILES_DB_SEPARATOR

	multiVers := ""
	if app.MultiVersion {
		multiVers = config.DECORATE_BOOL_FLAG_SYMBOL
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
	strArray := strings.Split(str, config.FILES_DB_SEPARATOR)

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
	if strArray[5] == config.FILES_DB_BOOL_FLAG_SYMBOL {
		app.MultiVersion = true
	}
	app.Family = strArray[6]
	return app
}
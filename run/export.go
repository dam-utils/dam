package run

import (
	"dam/config"
	"dam/driver/db"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"os"
)

func Export(arg string) {
	flag.ValidateProjectDirectory(arg)

	fs.Touch(arg)

	f, err := os.OpenFile(arg, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open apps file '%s' with error: %s", arg, err.Error())
	}

	apps := db.ADriver.GetApps()
	for _, app := range apps {
		newLine := app.ImageName+config.EXPORT_APP_SEPARATOR+app.ImageVersion+"\n"
		_, err := f.WriteString(newLine)
		if err != nil {
			logger.Fatal("Cannot write to apps file '%s' with error: %s", arg, err.Error())
		}
	}

	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync apps file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from apps file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}

	logger.Info("Export file save to '%s'", arg)
}

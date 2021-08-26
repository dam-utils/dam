package run

import (
	"dam/driver/conf/option"
	"os"

	"dam/driver/db"
	"dam/driver/engine"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/run/internal"
)

type ExportSettings struct {
	All bool
}

var ExportFlags = new(ExportSettings)

func Export(arg string) {
	flag.ValidateFilePath(arg)
	logger.Debug("Flags validated with success")

	absPath := fs.GetAbsolutePath(arg)

	if !ExportFlags.All {
		exportAppsListToFile(absPath)
		logger.Success("Export app list to file '%s'", absPath)
	} else {
		tmpDir := absPath + "_tmp"
		fs.MkDir(tmpDir)
		defer fs.Remove(tmpDir)

		logger.Debug("Exporting images file to tmp dir ...")
		exportAppsListToFile(tmpDir + string(os.PathSeparator) + option.Config.Export.GetAppsFileName())
		logger.Debug("Exporting docker images to tmp dir ...")
		exportImagesToDir(tmpDir)

		logger.Debug("Creating general apps archive ...")
		fs.Gzip(tmpDir, absPath, true)

		logger.Success("Export app list to apps archive '%s'", absPath)
	}
}

func exportImagesToDir(tmpDir string) {
	for _, app := range db.ADriver.GetApps() {
		tmpFilePath := tmpDir + string(os.PathSeparator) + option.Config.Save.GetTmpFilePostfix()
		tag := internal.GetPrefixRepo() + app.ImageName + ":" + app.ImageVersion
		logger.Info("Preparing image %s ...", tag)

		imageId := engine.VDriver.GetImageID(tag)
		if imageId == "" {
			logger.Fatal("Image with tag '%s' not exist in the system", tag)
		}
		engine.VDriver.SaveImage(imageId, tmpFilePath)

		modifyManifest(tmpFilePath, tag)
		resultPath := tmpDir +
			string(os.PathSeparator) +
			app.ImageName +
			option.Config.Save.GetFileSeparator() +
			app.ImageVersion +
			option.Config.Save.GetOptionalSeparator() +
			fs.HashFileCRC32(tmpFilePath) +
			option.Config.Save.GetFileSeparator() +
			fs.FileSize(tmpFilePath) +
			option.Config.Save.GetFilePostfix()
		fs.MoveFile(tmpFilePath, resultPath)
	}
}

func exportAppsListToFile(path string) {
	fs.Touch(path)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open apps file '%s' with error: %s", path, err)
	}

	logger.Debug("Getting apps ...")
	apps := db.ADriver.GetApps()
	for _, app := range apps {
		newLine := app.ImageName + option.Config.Export.GetAppSeparator() + app.ImageVersion + "\n"
		_, err := f.WriteString(newLine)
		if err != nil {
			logger.Fatal("Cannot write to apps file '%s' with error: %s", path, err)
		}
	}

	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync apps file '%s' with error: %s", path, err)
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from apps file '%s' with error: %s", path, err)
	}
}

package run

import (
	"os"
	"path"

	"dam/driver/conf/option"
	"dam/driver/engine"
	fs "dam/driver/filesystem"
	"dam/driver/filesystem/manifest"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/run/internal"
	"dam/run/internal/archive/app_name"
)

type SaveSettings struct {
	FilePath string
}

var SaveFlags = new(SaveSettings)

func Save(appFullName string) {
	var filePath string

	flag.ValidateAppPlusVersion(appFullName)
	logger.Debug("Flags validated with success")

	logger.Debug("Parsing tag ...")
	_, name, version := internal.SplitTag(appFullName)

	logger.Debug("Getting archive path ...")
	// TODO refactoring
	if SaveFlags.FilePath != "" {
		flag.ValidateFilePath(SaveFlags.FilePath)
		filePath = SaveFlags.FilePath

		logger.Debug("Saving archive ...")
		imageId := engine.VDriver.GetImageID(internal.GetPrefixRepo() + appFullName)
		if imageId == "" {
			logger.Fatal("Image with tag '%s' not exist in the system", internal.GetPrefixRepo()+appFullName)
		}
		engine.VDriver.SaveImage(imageId, filePath)

		logger.Debug("Preparing manifest ...")
		modifyManifest(filePath, appFullName)

		logger.Success("Created '%s' file.", filePath)
	} else {
		nameInfo := app_name.NewInfo()
		nameInfo.SetAppName(name)
		nameInfo.SetAppVersion(version)

		filePath = path.Join(fs.GetCurrentDir(), nameInfo.TempNameToString())

		logger.Debug("Saving archive ...")
		imageId := engine.VDriver.GetImageID(internal.GetPrefixRepo() + appFullName)
		if imageId == "" {
			logger.Fatal("Image with tag '%s' not exist in the system", internal.GetPrefixRepo()+appFullName)
		}
		engine.VDriver.SaveImage(imageId, filePath)

		logger.Debug("Preparing manifest ...")
		modifyManifest(filePath, appFullName)

		logger.Debug("Releasing archive ...")
		nameInfo.SetHash(fs.HashFileCRC32(filePath))
		nameInfo.SetSize(fs.FileSize(filePath))

		resultPath := path.Join(fs.GetCurrentDir(), nameInfo.FullNameToString())
		fs.MoveFile(filePath, resultPath)

		logger.Success("Created '%s' file.", resultPath)
	}
}

func modifyManifest(filePath, appFullName string) {
	dir := fs.Untar(filePath)
	manifestFile := dir + string(os.PathSeparator) + option.Config.Save.GetManifestFile()

	manifest.ModifyRepoTags(manifestFile, appFullName)
	fs.Remove(filePath)
	fs.Gzip(dir, filePath, false)
	fs.Remove(dir)
}

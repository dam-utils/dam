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

type saveArgType int
const (
	unknownSave saveArgType = iota
	appSave
	tagSave
)

func Save(arg string) {
	saveType := getSaveTypeByArg(arg)

	logger.Debug("Validating docker image with type '%v' ...", saveType)
	validateSaveArg(arg, saveType)
	logger.Debug("Flags validated with success")

	logger.Debug("Parsing tag ...")
	var imageTag string
	switch saveType {
	case appSave:
		imageTag = internal.GetPrefixRepo() + arg
	case tagSave:
		imageTag = arg
	}

	logger.Debug("Getting archive path ...")
	if SaveFlags.FilePath != "" {
		filePath := SaveFlags.FilePath
		flag.ValidateFilePath(filePath)

		logger.Debug("Saving archive ...")
		imageId := engine.VDriver.GetImageID(imageTag)
		if imageId == "" {
			logger.Fatal("Image with tag '%s' not exist in the system", imageTag)
		}
		engine.VDriver.SaveImage(imageId, filePath)

		logger.Debug("Preparing manifest ...")
		modifyManifest(filePath, imageTag)

		logger.Success("Created '%s' file.", filePath)
	} else {
		logger.Debug("Preparing tag ...")
		_, name, version := internal.SplitTag(arg)
		nameInfo := app_name.NewInfo()
		nameInfo.SetAppName(name)
		nameInfo.SetAppVersion(version)

		filePath := path.Join(fs.GetCurrentDir(), nameInfo.TempNameToString())

		logger.Debug("Saving archive to '%s' ...", filePath)
		imageId := engine.VDriver.GetImageID(imageTag)
		if imageId == "" {
			logger.Fatal("Image with tag '%s' not exist in the system", imageTag)
		}
		engine.VDriver.SaveImage(imageId, filePath)

		logger.Debug("Preparing manifest ...")
		modifyManifest(filePath, imageTag)

		logger.Debug("Releasing archive ...")
		nameInfo.SetHash(fs.HashFileCRC32(filePath))
		nameInfo.SetSize(fs.FileSize(filePath))

		resultPath := path.Join(fs.GetCurrentDir(), nameInfo.FullNameToString())
		fs.MoveFile(filePath, resultPath)

		logger.Success("Created '%s' file.", resultPath)
	}
}

func validateSaveArg(arg string, saveType saveArgType) {
	switch saveType {
	case appSave:
		flag.ValidateAppPlusVersion(arg)
	case tagSave:
		flag.ValidateTag(arg)
	default:
		logger.Fatal("Unknown argument '%s' for command 'save'. See '%s help save'", arg, option.Config.Global.GetProjectName())
	}
}

func getSaveTypeByArg(arg string) saveArgType {
	repo, name, version := internal.SplitTag(arg)
	logger.Debug("Split tag '%s' as Repo='%s', Name='%s' and Version='%s'", arg, repo, name, version)
	if repo != "" && name != "" && version != "" {
		return tagSave
	}

	if name != "" && version != "" {
		return appSave
	}

	return unknownSave
}

func modifyManifest(filePath, appFullName string) {
	dir := fs.Untar(filePath)
	manifestFile := dir + string(os.PathSeparator) + option.Config.Save.GetManifestFile()

	manifest.ModifyRepoTags(manifestFile, appFullName)
	fs.Remove(filePath)
	fs.Gzip(dir, filePath, false)
	fs.Remove(dir)
}

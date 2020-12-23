package run

import (
	"os"
	"path/filepath"

	"dam/config"
	"dam/driver/decorate"
	"dam/driver/engine"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/run/internal"
)

func InfoApp(tag string) {
	validateTag(tag)
	logger.Debug("Flags validated with success")

	logger.Debug("Getting meta ...")
	tmpDir := internal.PrepareTmpMetaPath(config.TMP_META_PATH)
	defer fs.Remove(tmpDir)

	containerId := engine.VDriver.ContainerCreate(tag, "")
	engine.VDriver.CopyFromContainer(containerId, string(os.PathSeparator)+config.META_DIR_NAME, tmpDir)
	engine.VDriver.ContainerRemove(containerId)

	logger.Debug("Printing description ...")
	decorate.Println()
	decorate.PrintDescription(filepath.Join(tmpDir, config.META_DIR_NAME, config.DESCRIPTION_FILE_NAME))
	decorate.Println()

	logger.Debug("Printing family label ...")
	family := internal.GetFamily(tag)
	decorate.PrintLabel(config.APP_FAMILY_ENV, family)

	logger.Debug("Printing multiversion label ...")
	imageId := engine.VDriver.GetImageID(tag)
	multiVersion, _ := engine.VDriver.GetImageLabel(imageId, config.APP_MULTIVERSION_ENV)
	if multiVersion != config.MULTIVERSION_TRUE_FLAG {
		multiVersion = config.MULTIVERSION_FALSE_FLAG
	}
	decorate.PrintLabel(config.APP_MULTIVERSION_ENV, multiVersion)
	decorate.Println()
}

func validateTag(tag string) {
	_, name, version := internal.SplitTag(tag)
	logger.Debug("Validating tag '%s' ...", tag)
	flag.ValidateAppName(name)
	flag.ValidateAppVersion(version)
}
package run

import (
	"os"
	"path/filepath"

	"dam/driver/conf/option"
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
	tmpDir := internal.PrepareTmpMetaPath(option.Config.FileSystem.GetTmpMetaPath())
	defer fs.Remove(tmpDir)

	containerId := engine.VDriver.ContainerCreate(tag, "")
	engine.VDriver.CopyFromContainer(containerId, string(os.PathSeparator)+option.Config.FileSystem.GetMetaDirName(), tmpDir)
	engine.VDriver.ContainerRemove(containerId)

	logger.Debug("Printing description ...")
	decorate.Println()
	decorate.PrintDescription(filepath.Join(tmpDir, option.Config.FileSystem.GetMetaDirName(), option.Config.FileSystem.GetDescriptionFileName()))
	decorate.Println()

	logger.Debug("Printing family label ...")
	family := internal.GetFamily(tag)
	decorate.PrintLabel(option.Config.ReservedEnvs.GetAppFamilyEnv(), family)

	logger.Debug("Printing multiversion label ...")
	imageId := engine.VDriver.GetImageID(tag)
	if imageId == "" {
		logger.Fatal("Image with tag '%s' not exist in the system", tag)
	}
	multiVersion, _ := engine.VDriver.GetImageLabel(imageId, option.Config.ReservedEnvs.GetAppMultiversionEnv())
	if multiVersion != option.Config.Multiversion.GetTrueFlag() {
		multiVersion = option.Config.Multiversion.GetFalseFlag()
	}
	decorate.PrintLabel(option.Config.ReservedEnvs.GetAppMultiversionEnv(), multiVersion)

	logger.Debug("Printing servers label ...")
	servers := internal.GetServersByTag(tag)
	decorate.PrintLabel(option.Config.ReservedEnvs.GetAppServersEnv(), servers)

	decorate.Println()
}

func validateTag(tag string) {
	_, name, version := internal.SplitTag(tag)
	logger.Debug("Validating tag '%s' ...", tag)
	flag.ValidateAppName(name)
	flag.ValidateAppVersion(version)
}

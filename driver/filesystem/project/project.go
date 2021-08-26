package project

import (
	"os"

	"dam/driver/conf/option"
	fs "dam/driver/filesystem"
	"dam/driver/filesystem/dockerfile"
	"dam/driver/logger"
)

func Prepare(path string) (string, string, string) {
	meta := path + string(os.PathSeparator) + option.Config.FileSystem.GetMetaDirName()
	if !fs.IsExistDir(meta) {
		logger.Fatal("Cannot find '%s' for path '%s'", option.Config.FileSystem.GetMetaDirName(), meta)
	}

	dockerFile := path + string(os.PathSeparator) + option.Config.FileSystem.GetDockerfileName()
	if !fs.IsExistFile(dockerFile) {
		logger.Fatal("Cannot find '%s' for path '%s'", option.Config.FileSystem.GetDockerfileName(), dockerFile)
	}

	install := meta + string(os.PathSeparator) + option.Config.FileSystem.GetInstallFileName()
	if !fs.IsExistFile(install) {
		if !fs.IsExistFile(install + option.Config.FileSystem.GetExpandMetaFile()) {
			logger.Fatal("Cannot find  '%s' or '%s%s' files in meta directory", install, install, option.Config.FileSystem.GetExpandMetaFile())
		}
	}

	uninstall := meta + string(os.PathSeparator) + option.Config.FileSystem.GetUninstallFileName()
	if !fs.IsExistFile(uninstall) {
		if !fs.IsExistFile(uninstall + option.Config.FileSystem.GetExpandMetaFile()) {
			logger.Fatal("Cannot find '%s' or '%s%s' files in meta directory", uninstall, uninstall, option.Config.FileSystem.GetExpandMetaFile())
		}
	}

	if !dockerfile.IsCopyMeta(dockerFile) {
		logger.Fatal("Not found COPY or ADD .. /meta command in Dockerfile '%s'", dockerFile)
	}

	if !dockerfile.IsFamily(dockerFile) {
		logger.Warn("Not found label 'FAMILY' in Dockerfile '%s'", dockerFile)
	}

	return meta, dockerFile, path + string(os.PathSeparator) + option.Config.FileSystem.GetEnvFileName()
}

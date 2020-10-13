package project

import (
	"os"

	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/filesystem/dockerfile"
	"dam/driver/logger"
)

func Prepare(path string)(string, string, string){
	meta := path+string(os.PathSeparator)+config.META_DIR_NAME
	if !fs.IsExistDir(meta) {
		logger.Fatal("Cannot found 'meta' for path '%s'", meta)
	}

	dockerFile := path+string(os.PathSeparator)+config.DOCKERFILE_NAME
	if !fs.IsExistFile(dockerFile) {
		logger.Fatal("Cannot found 'meta' for path '%s'", dockerFile)
	}

	install := meta+string(os.PathSeparator)+config.INSTALL_FILE_NAME
	if !fs.IsExistFile(install) {
		if !fs.IsExistFile(install+config.EXPAND_META_FILE) {
			logger.Fatal("Cannot found '%s' or '%s%s' files in meta directory", install, install, config.EXPAND_META_FILE)
		}
	}

	uninstall := meta+string(os.PathSeparator)+config.UNINSTALL_FILE_NAME
	if !fs.IsExistFile(uninstall) {
		if !fs.IsExistFile(uninstall+config.EXPAND_META_FILE) {
			logger.Fatal("Cannot found '%s' or '%s%s' files in meta directory", uninstall, uninstall, config.EXPAND_META_FILE)
		}
	}

	if !dockerfile.IsCopyMeta(dockerFile) {
		logger.Fatal("Not found COPY or ADD .. /meta command in Dockerfile '%s'", dockerFile)
	}

	if !dockerfile.IsFamily(dockerFile) {
		logger.Warn("Not found label 'FAMILY' in Dockerfile '%s'", dockerFile)
	}

	return meta, dockerFile, path+string(os.PathSeparator)+config.ENV_FILE_NAME
}

func ValidateTag(tag string) {
	// TODO
}


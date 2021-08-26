package option

import "dam/config"

type FileSystem struct {}

func (o *FileSystem) GetMetaDirName() string {
	return config.FS_META_DIR_NAME
}

func (o *FileSystem) GetDockerfileName() string {
	return config.FS_DOCKERFILE_NAME
}

func (o *FileSystem) GetEnvFileName() string {
	return config.FS_ENV_FILE_NAME
}

func (o *FileSystem) GetInstallFileName() string {
	return config.FS_INSTALL_FILE_NAME
}

func (o *FileSystem) GetUninstallFileName() string {
	return config.FS_UNINSTALL_FILE_NAME
}

func (o *FileSystem) GetDescriptionFileName() string {
	return config.FS_DESCRIPTION_FILE_NAME
}

func (o *FileSystem) GetExpandMetaFile() string {
	return config.FS_EXPAND_META_FILE
}

func (o *FileSystem) GetTmpMetaPath() string {
	return config.FS_TMP_META_PATH
}

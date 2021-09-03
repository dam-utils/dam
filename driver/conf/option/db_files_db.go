package option

import (
	"os"
	"path/filepath"
	"strconv"

	"dam/config"
)

type FilesDB struct {}

func (o *FilesDB) GetSeparator() string {
	return config.FILES_DB_SEPARATOR
}

func (o *FilesDB) GetBoolFlagSymbol() string {
	return config.FILES_DB_BOOL_FLAG_SYMBOL
}

func (o *FilesDB) GetFilesPermissions() os.FileMode {
	u64, err := strconv.ParseUint(config.FILES_DB_FILES_PERMISSIONS, 0, 32)
	if err != nil {
		printFatal("Config option 'FILES_DB_FILES_PERMISSIONS' is not valid for permission file mask: %s", err)
	}

	return os.FileMode(u64)
}

func (o *FilesDB) GetReposFilename() string {
	return getFullPath(config.FILES_DB_REPOS_FILENAME)
}

func (o *FilesDB) GetAppsFilename() string {
	return getFullPath(config.FILES_DB_APPS_FILENAME)
}

func (o *FilesDB) GetTmp() string {
	return getFullPath(config.FILES_DB_TMP)
}

func getFullPath(fullPath string) string {
	if config.FILES_DB_USE_USER_CACHE_DIR {
		dir, err := os.UserCacheDir()
		if err != nil {
			printFatal("Cannot get the user cache directory: %s", err)
		}
		fullPath = filepath.Join(dir, fullPath)
	}
	return fullPath
}


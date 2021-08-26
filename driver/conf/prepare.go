package conf

import (
	"dam/driver/conf/option"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
)

func Prepare() {
	switch option.Config.DB.GetType() {
	case "files":
		reposDir := fs.GetDir(option.Config.FilesDB.GetReposFilename())
		if !fs.IsExistDir(reposDir) {
			fs.MkDir(fs.GetDir(option.Config.FilesDB.GetReposFilename()))
		}
		fs.Touch(option.Config.FilesDB.GetReposFilename())

		appDir := fs.GetDir(option.Config.FilesDB.GetAppsFilename())
		if !fs.IsExistDir(appDir){
			fs.MkDir(fs.GetDir(option.Config.FilesDB.GetAppsFilename()))
		}
		fs.Touch(option.Config.FilesDB.GetAppsFilename())
	default:
		logger.Fatal("Cannot supported db '%s'", option.Config.DB.GetType())
	}
}

package internal

import (
	"dam/driver/conf/option"
	fs "dam/driver/filesystem"
)

func PrepareTmpMetaPath(meta string) string {
	path := fs.GetAbsolutePath(meta)
	fs.Remove(path)
	return path
}

func BoolToString(b bool) string {
	if b {
		return option.Config.Multiversion.GetTrueFlag()
	}
	return option.Config.Multiversion.GetFalseFlag()
}

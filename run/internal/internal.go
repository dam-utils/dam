package internal

import (
	"dam/config"
	fs "dam/driver/filesystem"
)

func PrepareTmpMetaPath(meta string) string {
	path := fs.GetAbsolutePath(meta)
	fs.Remove(path)
	return path
}

func BoolToString(b bool) string {
	if b {
		return config.MULTIVERSION_TRUE_FLAG
	}
	return config.MULTIVERSION_FALSE_FLAG
}

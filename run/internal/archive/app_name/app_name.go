package app_name

import (
	"fmt"
	"regexp"
	"strings"

	"dam/driver/conf/option"
)

type info struct {
	appName    string
	appVersion string
	hash string
	size string
}

func NewInfo() *info {
	return &info{
		appName:    option.Config.ReservedEnvs.GetDefaultAppName(),
		appVersion: option.Config.ReservedEnvs.GetDefaultAppVersion(),
		hash:       "0",
		size:       "0",
	}
}

func (i *info) FromString(str string) error {
	re := regexp.MustCompile(getRegexpMask())
	if !re.MatchString(str) {
		return fmt.Errorf("cannot match string by mask '%s'", getRegexpMask())
	}

	result1 := strings.TrimRight(str, option.Config.Save.GetFilePostfix())
	arrWithSize := strings.Split(result1, option.Config.Save.GetFileSeparator())
	i.size = arrWithSize[len(arrWithSize)-1]
	if i.size == "" {
		return fmt.Errorf("size is empty")
	}
	result2 := strings.TrimRight(result1, i.size)
	result3 := strings.TrimRight(result2, option.Config.Save.GetFileSeparator())
	arrWithHash := strings.Split(result3, option.Config.Save.GetOptionalSeparator())
	i.hash = arrWithHash[len(arrWithHash)-1]
	if i.hash == "" {
		return fmt.Errorf("hash is empty")
	}
	result4 := strings.TrimRight(result3, i.hash)
	result5 := strings.TrimRight(result4, option.Config.Save.GetOptionalSeparator())
	arrWithVersion := strings.Split(result5, option.Config.Save.GetFileSeparator())
	i.appVersion = arrWithVersion[len(arrWithVersion)-1]
	if i.appVersion == "" {
		return fmt.Errorf("app version is empty")
	}
	result6 := strings.TrimRight(result5, i.appVersion)
	i.appName = strings.TrimRight(result6, option.Config.Save.GetFileSeparator())
	if i.appName == "" {
		return fmt.Errorf("app name is empty")
	}
	return nil
}

func (i *info) FullNameToString() string {
	return i.appName +
		option.Config.Save.GetFileSeparator() +
		i.appVersion +
		option.Config.Save.GetOptionalSeparator() +
		i.hash +
		option.Config.Save.GetFileSeparator() +
		i.size +
		option.Config.Save.GetFilePostfix()
}

func (i *info) TempNameToString() string {
	return i.appName +
		option.Config.Save.GetFileSeparator() +
		i.appVersion +
		option.Config.Save.GetTmpFilePostfix()
}

func (i *info) SetAppName(str string) {
	i.appName = str
}

func (i *info) SetAppVersion(str string) {
	i.appVersion = str
}

func (i *info) SetHash(str string) {
	i.hash = str
}

func (i *info) Hash() string {
	return i.hash
}

func (i *info) SetSize(str string) {
	i.size = str
}

func (i *info) Size() string {
	return i.size
}

func getRegexpMask() string {
	return ".*" +
		option.Config.Save.GetFileSeparator() +
		".*" +
		option.Config.Save.GetOptionalSeparator() +
		".*" +
		option.Config.Save.GetFileSeparator() +
		".*" +
		option.Config.Save.GetFilePostfix() +
		"\\b"
}

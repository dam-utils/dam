package app_name

import (
	"fmt"
	"regexp"
	"strings"

	"dam/driver/conf/option"
)

type info struct {
	AppName    string
	AppVersion string
	Hash       string
	Size       string
}

func NewInfo() *info {
	return &info{
		AppName:    option.Config.ReservedEnvs.GetDefaultAppName(),
		AppVersion: option.Config.ReservedEnvs.GetDefaultAppVersion(),
		Hash:       "0",
		Size:       "0",
	}
}

func (i *info) FromString(str string) error {
	re := regexp.MustCompile(getRegexpMask())
	if !re.MatchString(str) {
		return fmt.Errorf("cannot match string by mask '%s'", getRegexpMask())
	}

	result1 := strings.TrimRight(str, option.Config.Save.GetFilePostfix())
	arrWithSize := strings.Split(result1, option.Config.Save.GetFileSeparator())
	i.Size = arrWithSize[len(arrWithSize)-1]
	if i.Size == "" {
		return fmt.Errorf("size is empty")
	}
	result2 := strings.TrimRight(result1, i.Size)
	result3 := strings.TrimRight(result2, option.Config.Save.GetFileSeparator())
	arrWithHash := strings.Split(result3, option.Config.Save.GetOptionalSeparator())
	i.Hash = arrWithHash[len(arrWithHash)-1]
	if i.Hash == "" {
		return fmt.Errorf("hash is empty")
	}
	result4 := strings.TrimRight(result3, i.Hash)
	result5 := strings.TrimRight(result4, option.Config.Save.GetOptionalSeparator())
	arrWithVersion := strings.Split(result5, option.Config.Save.GetFileSeparator())
	i.AppVersion = arrWithVersion[len(arrWithVersion)-1]
	if i.AppVersion == "" {
		return fmt.Errorf("app version is empty")
	}
	result6 := strings.TrimRight(result5, i.AppVersion)
	i.AppName = strings.TrimRight(result6, option.Config.Save.GetFileSeparator())
	if i.AppName == "" {
		return fmt.Errorf("app name is empty")
	}
	return nil
}

func (i *info) FullNameToString() string {
	return i.AppName +
		option.Config.Save.GetFileSeparator() +
		i.AppVersion +
		option.Config.Save.GetOptionalSeparator() +
		i.Hash +
		option.Config.Save.GetFileSeparator() +
		i.Size +
		option.Config.Save.GetFilePostfix()
}

func (i *info) TempNameToString() string {
	return i.AppName +
		option.Config.Save.GetFileSeparator() +
		i.AppVersion +
		option.Config.Save.GetTmpFilePostfix()
}

func (i *info) SetAppName(str string) {
	i.AppName = str
}

func (i *info) SetAppVersion(str string) {
	i.AppVersion = str
}

func (i *info) SetHash(str string) {
	i.Hash = str
}

func (i *info) SetSize(str string) {
	i.Size = str
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

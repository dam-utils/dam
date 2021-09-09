package app_name

import (
	"dam/driver/logger"
	"fmt"
	"regexp"
	"strings"

	"dam/driver/conf/option"
)

type info struct {
	appName    string
	appVersion string
	hash       string
}

func NewInfo() *info {
	return &info{
		appName:    option.Config.ReservedEnvs.GetDefaultAppName(),
		appVersion: option.Config.ReservedEnvs.GetDefaultAppVersion(),
		hash:       "0",
	}
}

func (i *info) FromString(str string) error {
	err := checkDoubleSymbol(str, option.Config.Save.GetOptionalSeparator())
	if err != nil {
		return err
	}

	err = checkDoubleSymbol(str, option.Config.Save.GetFileSeparator())
	if err != nil {
		return err
	}

	re := regexp.MustCompile(getRegexpMask())
	if !re.MatchString(str) {
		return fmt.Errorf("cannot match string by mask '%s'", getRegexpMask())
	}

	result1 := strings.TrimRight(str, option.Config.Save.GetFilePostfix())
	arrWithHash := strings.Split(result1, option.Config.Save.GetOptionalSeparator())
	i.hash = arrWithHash[len(arrWithHash)-1]
	logger.Debug("Split prefix filename with hash: '%s' and hash '%s'\n", arrWithHash, i.hash)
	if i.hash == "" {
		return fmt.Errorf("hash is empty")
	}
	result3 := strings.TrimRight(result1, i.hash)
	result4 := strings.TrimRight(result3, option.Config.Save.GetOptionalSeparator())
	arrWithVersion := strings.Split(result4, option.Config.Save.GetFileSeparator())
	i.appVersion = arrWithVersion[len(arrWithVersion)-1]
	if i.appVersion == "" {
		return fmt.Errorf("app version is empty")
	}
	result5 := strings.TrimRight(result4, i.appVersion)
	i.appName = strings.TrimRight(result5, option.Config.Save.GetFileSeparator())
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

func getRegexpMask() string {
	return ".*" +
		option.Config.Save.GetFileSeparator() +
		".*" +
		option.Config.Save.GetOptionalSeparator() +
		".*" +
		option.Config.Save.GetFilePostfix() +
		"\\b"
}

func checkDoubleSymbol(str, symbol string) error {
	logger.Debug("Check filename for double symbol '%s'", symbol)
	re := regexp.MustCompile("\\" +	symbol + "\\" +	symbol)
	if re.MatchString(str) {
		return fmt.Errorf("string has two consecutive '%s'", symbol)
	}
	return nil
}
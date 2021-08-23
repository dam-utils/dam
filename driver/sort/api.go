package sort

import (
	"dam/driver/conf/option"
	"dam/driver/logger"
)

func SortAppNames(names *[]string) *[]string {
	switch option.Config.Sort.GetAppType() {
	case "alphabetic":
		return AlphabeticSort(names)
	default:
		logger.Fatal("Config option SORT_APP_TYPE='%s' not valid. Sorting type is bad", option.Config.Sort.GetAppType())
	}
	return nil
}

func SortVersions(vers *[]string) *[]string {
	switch option.Config.Sort.GetVersionType() {
	case "semantic_version":
		return SemanticVersionSort(vers)
	default:
		logger.Fatal("Config option SORT_VERSION_TYPE='%s' not valid. Sorting type is bad", option.Config.Sort.GetVersionType())
	}
	return nil
}
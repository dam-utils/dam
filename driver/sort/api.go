package sort

import (
	"dam/config"
	"dam/driver/logger"
)

func SortAppNames(names *[]string) *[]string {
	switch config.SORT_APP_TYPE {
	case "alphabetic":
		return AlphabeticSort(names)
	default:
		logger.Fatal("Config option SORT_APP_TYPE='%s' not valid. Sorting type is bad", config.SORT_APP_TYPE)
	}
	return nil
}

func SortVersions(vers *[]string) *[]string {
	switch config.SORT_VERSION_TYPE {
	case "semantic_version":
		return SemanticVersionSort(vers)
	default:
		logger.Fatal("Config option SORT_VERSION_TYPE='%s' not valid. Sorting type is bad", config.SORT_VERSION_TYPE)
	}
	return nil
}
package run_test

import (
	"log"
	"os"
	"testing"

	"dam/config"
	"dam/driver/conf/option"
	"dam/driver/db"
	"dam/driver/logger"
)

func init() {
	log.SetFlags(0)
	logger.DebugMode = false
	db.Init()
}

func setDefaultConfig() {
	switch option.Config.DB.GetType() {
	case "files":
		config.DECORATE_MAX_DISPLAY_WIDTH = 100
		config.FILES_DB_USE_USER_CACHE_DIR = false
		config.FILES_DB_REPOS_FILENAME = "Repos"
		config.FILES_DB_APPS_FILENAME = "Apps"
		config.FILES_DB_TMP = ".db"
		config.FILES_DB_SEPARATOR = "|"
	default:
		logger.Fatal("Cannot supported db '%s'", option.Config.DB.GetType())
	}
	config.DECORATE_RAW_SEPARATOR = "|"
	config.OFFICIAL_REGISTRY_URL = "https://registry-1.docker.io/"
}

func dropTestDB(t *testing.T) {
	for _, path := range [...]string{config.FILES_DB_REPOS_FILENAME, config.FILES_DB_APPS_FILENAME, config.FILES_DB_TMP} {
		_, err := os.Stat(path)
		if err == nil {
			err = os.Remove(path)
			if err != nil {
				t.Fail()
			} else {
				t.Log("FilePath '" + path + "' was removed")
			}
		} else {
			t.Log("Skip removing '" + path + "' file")
		}
	}
}

func printFailInfo(t *testing.T, expResult string, result string) {
	if result != expResult {
		t.Log("Expected result:")
		t.Log("`" + expResult + "`")
		t.Log("Result:")
		t.Log("`" + result + "`")
		t.Fail()
	}
}

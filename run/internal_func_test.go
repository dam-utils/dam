package run_test

import (
	"dam/driver/logger"
	"log"
	"os"
	"testing"

	"dam/config"
	"dam/driver/db"
)

func init(){
	log.SetFlags(0)
	db.Init()
}

func setDefaultConfig() {
	switch config.DB_TYPE {
	case "files":
		config.DECORATE_MAX_DISPLAY_WIDTH = 100
		config.FILES_DB_REPOS = "Repos"
		config.FILES_DB_APPS = "Apps"
		config.FILES_DB_TMP = ".db"
		config.FILES_DB_SEPARATOR = "|"
	default:
		logger.Fatal("Cannot supported db '%s'", config.DB_TYPE)
	}
	config.DECORATE_RAW_SEPARATOR = "|"
	config.OFFICIAL_REGISTRY_URL = "https://registry-1.docker.io/"
}

func dropTestDB(t *testing.T) {
	for _, path := range [...]string{config.FILES_DB_REPOS, config.FILES_DB_REPOS, config.FILES_DB_TMP}{
		_, err := os.Stat(path)
		if err == nil {
			err = os.Remove(path)
			if err != nil{
				t.Fail()
			} else {
				t.Log("FilePath '"+path+"' was removed")
			}
		} else {
			t.Log("Skip removing '"+path+"' file")
		}
	}
}

func printFailInfo(t *testing.T, expResult string, result string) {
	if result != expResult {
		t.Log("Expected result:")
		t.Log("`"+expResult+"`")
		t.Log("Result:")
		t.Log("`"+result+"`")
		t.Fail()
	}
}
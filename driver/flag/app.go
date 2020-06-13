package flag

import (
	"dam/driver/logger"
	"dam/driver/validate"
)

func ValidateAppName(s string) {
	err := validate.CheckAppName(s)
	if err != nil {
		logger.Error(err.Error())
		logger.Fatal("App name flag is not valid.")
	}
}

func ValidateAppVersion(s string) {
	err := validate.CheckVersion(s)
	if err != nil {
		logger.Error(err.Error())
		logger.Fatal("App version flag is not valid.")
	}
}

func ValidateAppPlusVersion(s string) {
	err := validate.CheckApp(s)
	if err != nil {
		logger.Error(err.Error())
		logger.Fatal("<app>:<version> flag is not valid.")
	}
}

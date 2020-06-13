package flag

import (
	"dam/driver/logger"
	"dam/driver/validate"
)

func ValidateProjectDirectory(s string) {
	err := validate.ProjectDir(s)
	if err != nil {
		logger.Error(err.Error())
		logger.Fatal("Project directory flag is not valid.")
	}
}

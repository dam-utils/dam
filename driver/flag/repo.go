package flag

import (
	"dam/driver/logger"
	"dam/driver/validate"
)

func ValidateRepoName(s string) {
	err := validate.CheckRepoName(s)
	if err != nil {
		logger.Fatal(err.Error())
		logger.Fatal("Repository name flag is not valid.")
	}
}

func ValidateRepoServer(s string) {
	err := validate.CheckServer(s)
	if err != nil {
		logger.Error(err.Error())
		logger.Fatal("Server url flag is not valid.")
	}
}

func ValidateRepoUsername(s string) {
	err := validate.CheckLogin(s)
	if err != nil {
		logger.Error(err.Error())
		logger.Fatal("Username flag is not valid.")
	}
}

func ValidateRepoPassword(s string) {
	err := validate.CheckPassword(s)
	if err != nil {
		logger.Error(err.Error())
		logger.Fatal("Password flag is not valid.")
	}
}

func ValidateRepoID(s string) {
	err := validate.CheckID(s)
	if err != nil {
		logger.Error(err.Error())
		logger.Fatal("ID flag is not valid.")
	}
}

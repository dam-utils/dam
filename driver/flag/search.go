package flag

import (
	"dam/driver/logger"
	"dam/driver/validate"
)

func ValidateSearchMask(s string) {
	err := validate.CheckMask(s)
	if err != nil {
		logger.Error(err.Error())
		logger.Fatal("Search mask flag is not valid.")
	}
}

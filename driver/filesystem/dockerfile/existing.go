package dockerfile

import (
	"bufio"
	"dam/driver/conf/option"
	"os"
	"regexp"

	"dam/driver/logger"
)

func IsCopyMeta(path string) bool {
	f, err := os.Open(path)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open Dockerfile with error: %s", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if regexp.MustCompile("COPY .* /meta").Match(scanner.Bytes()) {
			return true
		}

		if regexp.MustCompile("ADD .* /meta").Match(scanner.Bytes()) {
			return true
		}
	}
	return false
}

func IsFamily (path string) bool {
	f, err := os.Open(path)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open Dockerfile with error: %s", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		matchFamily, err := regexp.MatchString("LABEL "+option.Config.ReservedEnvs.GetAppFamilyEnv()+"=.* ", line)
		if err != nil {
			if matchFamily {
				return true
			}
		}
	}
	return false
}
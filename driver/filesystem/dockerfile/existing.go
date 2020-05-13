// Copyright 2020 The Docker Applications Manager Authors
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
package dockerfile

import (
	"bufio"
	"os"
	"regexp"

	"dam/config"
	"dam/driver/logger"
)

func IsCopyMeta(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		logger.Fatal("Cannot open Dockerfile with error: " + err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
			line := scanner.Text()
			matchCopy, err := regexp.MatchString("^COPY .* /meta$", line)
			if err != nil {
				if matchCopy {
					return true
				}
			}
			matchAdd, err := regexp.MatchString("^ADD .* /meta$", line)
			if err != nil{
				if matchAdd {
					return true
				}
			}
	}
	return false
}

func IsFamily (path string) bool {
	f, err := os.Open(path)
	if err != nil {
		logger.Fatal("Cannot open Dockerfile with error: " + err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		matchFamily, err := regexp.MatchString("^LABEL "+config.APP_FAMILY+"$", line)
		if err != nil {
			if matchFamily {
				return true
			}
		}
	}
	return false
}
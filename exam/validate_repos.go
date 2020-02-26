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
package exam

import (
	"fmt"
	"regexp"
	"strconv"
)

func CheckRepoName(name string) error {
	l := len(name)
	if l<3 && l>9 {
		return fmt.Errorf("Repository name is bad. It must have lenght 3-9 symbols")
	}

	regexPattern := "[A-Za-z0-9_]"
	matched, err := regexp.Match(regexPattern, []byte(name))
	if err != nil {
		return fmt.Errorf("Internal error. Cannot match regex patern '" +
			regexPattern + "' with registry name '" + name + "'")
	}
	if !matched {
		return fmt.Errorf("Repository name is bad. It must have only letters, numbers and '_'")
	}

	_, err = strconv.ParseInt(name, 10, 32)
	if err == nil {
		return fmt.Errorf("Repository name is bad. It cannot be a registry number (ID)")
	}

	return nil
}

func CheckServer(server string) error {
	l := len(server)
	if l>120 {
		return fmt.Errorf("Server option is bad. It must have lenght '<' or '=' 120 symbols")
	}

	if server == "" {
		return fmt.Errorf("Server name is not valid. It cannot be an empty string")
	}

	return nil
}

func CheckLogin(login string) error {
	l := len(login)
	if l>24 {
		return fmt.Errorf("Login option is bad. It must have lenght '<' or '=' 24 symbols")
	}
	return nil
}

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
package env

import (
	"bufio"
	"dam/driver/logger"
	"os"
	"strings"
)

func GetFileEnv(file string) map[string]string {
	var envMap = make(map[string]string)

	f, err := os.Open(file)
	if err != nil {
		logger.Warn("Cannot open env file with error: %s", err.Error())
		return envMap
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		eKey, eVar, ok := convertEnvFileLIne(scanner.Text())
		if ok {
			envMap[eKey] = eVar
		}
	}
	return envMap
}

func GetDockerFileEnv(file string) map[string]string {
	var envMap = make(map[string]string)

	f, err := os.Open(file)
	if err != nil {
		logger.Fatal("Cannot open docker file with error: %s", err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		eKey, eVar, ok := convertDockerFileLIne(scanner.Text())
		if ok {
			envMap[eKey] = eVar
		}
	}
	return envMap
}

func GetOSEnv(envPrefix string) map[string]string {
	var envMap = make(map[string]string)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			if strings.HasPrefix(pair[0], envPrefix) {
				envMap[pair[0]]=pair[1]
			}
		}
	}
	return envMap
}

// первый параметр функции в меньшем приоритете перед вторым при слиянии
func MergeEnvs(map1, map2 map[string]string) map[string]string {
	for key2, var2 := range map2 {
		map1[key2] = var2
	}
	return map1
}

func convertEnvFileLIne(line string) (string, string, bool){
	// Формат ENVIRONMENT файла:
	//- Не игнорируются пробелы и табуляции
	//- комментарии начинаются с символа '#'
	//- нельзя использовать переменную в несколько строчек
	//- нельзя делать комментарии в той же строке, что и переменная
	//- строчки без '=' игнорируются
	if strings.HasPrefix(line,"#") {
		return "","", false
	}
	splResult := strings.SplitN(line, "=", 2)
	if len(splResult) != 2 {
		return "","", false
	}
	return splResult[0], splResult[1], true
}

func convertDockerFileLIne(line string) (string, string, bool){
	// Формат Dockerfile:
	//- Пример `ENV FOO=foo`
	//- Не игнорируются пробелы и табуляции
	//- нельзя переменную в несколько строчек
	//- нельзя комментарии в той же строке, что и переменная
	//- разделителем имени переменной могут быть символы "=" и " "
	//- переменные без значения игнорируются
	if !strings.HasPrefix(line,"ENV "){
		return "","", false
	}

	newline := line[4:]
	splResult1 := strings.SplitN(newline, "=", 2)
	if len(splResult1) != 2 {
		splResult2 := strings.SplitN(newline, " ", 2)
		if len(splResult2) != 2 {
			return "", "", false
		}
		return splResult2[0], splResult2[1], true
	}
	return splResult1[0], splResult1[1], true
}


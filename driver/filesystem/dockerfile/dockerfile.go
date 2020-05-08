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
	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
)

// проверить, существует ли Dockerfile по флагу
// проверить, существует ли path по флагу + "Dockerfile "
// проверить, существует ли Dockerfile в текущем каталоге
func GetPath(filePath string, pwd string) string {
	if filePath != "" {
		if fs.FileIsExist(filePath) {
			if fs.GetBaseName(filePath) == config.DOCKERFILE_NAME {
				return filePath
			}
		}
		if fs.FileIsExist(filePath + "/" + config.DOCKERFILE_NAME) {
			return filePath + "/" + config.DOCKERFILE_NAME
		}
	}
	if fs.FileIsExist(pwd+"/"+config.DOCKERFILE_NAME){
		return pwd+"/"+config.DOCKERFILE_NAME
	}
	logger.Fatal("Cannot found Dockerfile path in command")
	// Never gonna happen
	return ""
}

func PrepareCopyMeta(dockerFile, metaDir, metaFlag string) {
	//Если есть флаг, с указанием на мету. Значит мета по нестандартному пути. Проверить:
	//- Если указано в Dockerfile копирование меты - ошибка
	//- Добавить копирование в dockerfile
	//Если нет флага:
	//- Если указано в Dockerfile копирование меты - ок
	//- Если не указано в Dockerfile копирование меты - указать

	if metaFlag == "" {
		if checkCopyMeta(dockerFile) {
			return
		}
	} else {
		if checkCopyMeta(dockerFile) {
			logger.Fatal("Found conflict. Dockerfile contain string 'COPY ... /meta' and flag contain path of meta dir.\nPlease select one project configuration.")
		}
	}
	sedCopyMeta(dockerFile, metaDir)
}

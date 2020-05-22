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
package config

var (
	//Global options // TODO fix it
	UTIL_NAME="dam"
	DISABLE_DEBUG=true //	It is recommended to enable this option in a production. Debug messages may be contain passwords
	//Decorating options
	COLLOR_ON=true
	DECORATE_MAX_DISPLAY_WIDTH =100
	DECORATE_RAW_SEPARATOR     ="|"
	DECORATE_BOOL_FLAG         ="*"
	//DB types
	DB_TYPE            ="files" // TODO create new dbs
	//FILES DB options
	FILES_DB_VALIDATE  =true
	FILES_DB_SEPARATOR ="|"
	FILES_DB_BOOL_FLAG ="*"
	FILES_DB_REPOS     ="src/examples/db/files/Repos"
	FILES_DB_APPS      ="src/examples/db/files/Apps"
	FILES_DB_TMP       ="src/examples/db/files/.db"
	//Repositories
	SEARCH_PROTOCOL_STRATEGY=[]string{"https","http"} // The order of the protocols is important.
	SEARCH_MAX_CONNECTS=1
	SEARCH_TIMEOUT_MS=1000
	OFFICIAL_REGISTRY_AUTH_URL="https://auth.docker.io/token?service=registry.docker.io"
	OFFICIAL_REGISTRY_URL="https://registry-1.docker.io/"
	OFFICIAL_REGISTRY_NAME="official" // TODO delete hardcode in tests
	//Docker
	DOCKER_API_VERSION="1.40"
	//Search
	OFFICIAL_REPO_SEARCH_APPS_LIMIT=100  // [1,100]
	INTERNAL_REPO_SEARCH_APPS_LIMIT=999
	//Sorting
	SORT_APP_TYPE="alphabetic" // TODO create new sorting
	SORT_VERSION_TYPE="semantic_version" // TODO create new sorting
	//Creating
	META_DIR_NAME="meta"
	DOCKERFILE_NAME="Dockerfile"
	ENV_FILE_NAME="ENVIRONMENT"
	INSTALL_FILE_NAME="install"
	UNINSTALL_FILE_NAME="uninstall"

	OS_ENV_PREFIX="DAM_"
	EXPAND_META_FILE=".exp"
	//Reserved ENVs
	APP_NAME_ENV="DAM_APP_NAME"
	DEF_APP_NAME="unknown"
	APP_VERS_ENV="DAM_APP_VERSION"
	DEF_APP_VERS="SNAPSHOT"
	APP_FAMILY="DAM_APP_FAMILY"
	//Install app
	TMP_META_PATH="./.tmp.meta"
)
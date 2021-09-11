package config

var (
	//Global options
	PROJECT_NAME    = "dam"
	PROJECT_VERSION = "1.2.1"

	//Decorating options
	COLOR_ON                   = true
	DECORATE_MAX_DISPLAY_WIDTH = 100
	DECORATE_RAW_SEPARATOR     = ";"
	DECORATE_BOOL_FLAG_SYMBOL  = "*"

	//DB types
	DB_TYPE = "files"

	//FILES DB options. Type: "files"
	FILES_DB_SEPARATOR         = ";"
	FILES_DB_BOOL_FLAG_SYMBOL  = "*"
	FILES_DB_FILES_PERMISSIONS = "0660"
	// If enabled this FilesDB option DAM will save db files to the user cache directory
	FILES_DB_USE_USER_CACHE_DIR = true
	FILES_DB_REPOS_FILENAME     = "dam/Repos"
	FILES_DB_APPS_FILENAME      = "dam/Apps"
	FILES_DB_TMP                = "dam/.db"

	//Virtualization type
	VIRTUALIZATION_TYPE = "docker"

	//Search
	SEARCH_PROTOCOL_STRATEGY        = "https, http" // The order of the protocols is important.
	SEARCH_MAX_CONNECTIONS          = 1
	SEARCH_TIMEOUT_MS               = 1000
	SEARCH_OFFICIAL_REPO_APPS_LIMIT = 100 // [1,100]
	SEARCH_INTERNAL_REPO_APPS_LIMIT = 999

	//OfficialRepo
	OFFICIAL_REGISTRY_AUTH_URL = "https://auth.docker.io/token?service=registry.docker.io"
	OFFICIAL_REGISTRY_URL      = "https://registry-1.docker.io/"
	OFFICIAL_REGISTRY_NAME     = "official" // TODO delete hardcode in tests

	//DefaultRepo
	UNKNOWN_REPO_NAME      = "~unknown~"
	UNKNOWN_REPO_ID        = 0
	NEW_REPO_PREFIX        = "auto"
	NEW_REPO_POSTFIX_LIMIT = 999
	LABEL_REPOS_SEPARATOR  = ","

	//Docker
	DOCKER_API_VERSION = "1.40"

	//Sort
	SORT_APP_TYPE     = "alphabetic"
	SORT_VERSION_TYPE = "semantic_version"

	//FileSystem
	FS_META_DIR_NAME         = "meta"
	FS_DOCKERFILE_NAME       = "Dockerfile"
	FS_ENV_FILE_NAME         = "ENVIRONMENT"
	FS_INSTALL_FILE_NAME     = "install"
	FS_UNINSTALL_FILE_NAME   = "uninstall"
	FS_DESCRIPTION_FILE_NAME = "DESCRIPTION"
	FS_EXPAND_META_FILE      = ".exp"
	FS_TMP_META_PATH         = "./.tmp.meta"

	//Multiversion
	MULTIVERSION_TRUE_FLAG  = "true"
	MULTIVERSION_FALSE_FLAG = "false"

	//Reserved Envs
	OS_ENV_PREFIX        = "DAM_"
	DEF_APP_NAME         = "unknown"
	DEF_APP_VERS         = "SNAPSHOT"
	APP_NAME_ENV         = "DAM_APP_NAME"
	APP_VERS_ENV         = "DAM_APP_VERSION"
	APP_FAMILY_ENV       = "DAM_APP_FAMILY"
	APP_MULTIVERSION_ENV = "DAM_APP_MULTIVERSION"
	APP_TAG_ENV          = "DAM_APP_TAG"
	APP_SERVERS_ENV      = "DAM_APP_SERVERS"

	//Export
	EXPORT_APP_SEPARATOR  = ":"
	EXPORT_APPS_FILE_NAME = "app_list"

	//Save
	SAVE_OPTIONAL_SEPARATOR        = "."
	SAVE_FILE_SEPARATOR            = "-"
	SAVE_TMP_FILE_POSTFIX          = ".dam.tmp"
	SAVE_FILE_POSTFIX              = ".dam"
	SAVE_POLYNOMIAL_CKSUM   uint32 = 0xedb88320
	SAVE_MANIFEST_FILE             = "manifest.json"
)

package config

var (
	//Global options
	PROJECT_NAME    = "dam"
	PROJECT_VERSION = "1.1.1"

	//Decorating options
	COLLOR_ON                  = true
	DECORATE_MAX_DISPLAY_WIDTH = 100
	DECORATE_RAW_SEPARATOR     = ";"
	DECORATE_BOOL_FLAG         = "*"
	//DB types
	DB_TYPE = "files" // TODO create new dbs
	//FILES DB options. Type: "files"
	// WARN: Use only absolutely files path
	FILES_DB_SEPARATOR = ";"
	FILES_DB_BOOL_FLAG = "*"
	FILES_DB_REPOS     = "/tmp/Repos"
	FILES_DB_APPS      = "/tmp/Apps"
	FILES_DB_TMP       = "/tmp/.db"
	// ContainerD type
	VIRTUALIZATION_TYPE = "docker"
	//Repositories
	SEARCH_PROTOCOL_STRATEGY   = []string{"https", "http"} // The order of the protocols is important.
	SEARCH_MAX_CONNECTS        = 1
	SEARCH_TIMEOUT_MS          = 1000
	OFFICIAL_REGISTRY_AUTH_URL = "https://auth.docker.io/token?service=registry.docker.io"
	OFFICIAL_REGISTRY_URL      = "https://registry-1.docker.io/"
	OFFICIAL_REGISTRY_NAME     = "official"                // TODO delete hardcode in tests
	UNKNOWN_REPO_NAME          = "~unknown~"
	NEW_REPO_PREFIX            = "auto"
	NEW_REPO_POSTFIX_LIMIT     = 999

	//Docker
	DOCKER_API_VERSION = "1.40"
	//Search
	OFFICIAL_REPO_SEARCH_APPS_LIMIT = 100 // [1,100]
	INTERNAL_REPO_SEARCH_APPS_LIMIT = 999
	//Sorting
	SORT_APP_TYPE     = "alphabetic"       // TODO create new sorting
	SORT_VERSION_TYPE = "semantic_version" // TODO create new sorting
	//Creating
	META_DIR_NAME         = "meta"
	DOCKERFILE_NAME       = "Dockerfile"
	ENV_FILE_NAME         = "ENVIRONMENT"
	INSTALL_FILE_NAME     = "install"
	UNINSTALL_FILE_NAME   = "uninstall"
	DESCRIPTION_FILE_NAME = "DESCRIPTION"
	OS_ENV_PREFIX         = "DAM_"
	EXPAND_META_FILE      = ".exp"
	LABEL_REPOS_SEPARATOR = ","
	//Multiversion
	MULTIVERSION_TRUE_FLAG  = "true"
	MULTIVERSION_FALSE_FLAG = "false"
	//Reserved ENVs
	APP_NAME_ENV         = "DAM_APP_NAME"
	DEF_APP_NAME         = "unknown"
	APP_VERS_ENV         = "DAM_APP_VERSION"
	DEF_APP_VERS         = "SNAPSHOT"
	APP_FAMILY_ENV       = "DAM_APP_FAMILY"
	APP_MULTIVERSION_ENV = "DAM_APP_MULTIVERSION"
	APP_TAG_ENV          = "DAM_APP_TAG"
	APP_SERVERS_ENV		 = "DAM_APP_SERVERS"
	//Install app
	TMP_META_PATH = "./.tmp.meta"
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

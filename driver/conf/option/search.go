package option

import (
	"strings"
	"time"

	"dam/config"
)

type Search struct {}

func (o *Search) GetProtocolStrategy() []string {
	result := strings.Split(config.SEARCH_PROTOCOL_STRATEGY, ",")
	if len(result) == 0 {
		printFatal("Not correct config option SEARCH_PROTOCOL_STRATEGY. Example: SEARCH_PROTOCOL_STRATEGY=\"https, http\"")
	}

	return result
}

func (o *Search) GetMaxConnections() int {
	return config.SEARCH_MAX_CONNECTIONS
}

func (o *Search) GetTimeoutMs() time.Duration {
	// TODO Проверить правильное ли это выражение
	return time.Duration(config.SEARCH_TIMEOUT_MS) * time.Millisecond
}

func (o *Search) GetOfficialRepoAppsLimit() int {
	return config.SEARCH_OFFICIAL_REPO_APPS_LIMIT
}

func (o *Search) GetInternalRepoAppsLimit() int {
	return config.SEARCH_INTERNAL_REPO_APPS_LIMIT
}
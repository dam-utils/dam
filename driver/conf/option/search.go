package option

import (
	"dam/config"
	"dam/driver/logger"
	"strings"
)

type Search struct {}

func (o *Search) GetProtocolStrategy() []string {
	result := strings.Split(config.SEARCH_PROTOCOL_STRATEGY, ",")
	if len(result) == 0 {
		logger.Fatal("Not correct config option SEARCH_PROTOCOL_STRATEGY. Example: SEARCH_PROTOCOL_STRATEGY=\"https, http\"")
	}

	return result
}

func (o *Search) GetMaxConnections() int {
	return config.SEARCH_MAX_CONNECTIONS
}

func (o *Search) GetTimeoutMs() int {
	return config.SEARCH_TIMEOUT_MS
}

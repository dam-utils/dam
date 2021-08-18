package apps

import (
	"dam/config"
	"dam/driver/logger"
	"os"
)

func (p *provider) connect() {
	var err error
	p.client, err = os.Open(config.FILES_DB_APPS_FILENAME)
	if err != nil {
		logger.Fatal("Cannot open db file '%s' with error: %s", config.FILES_DB_APPS_FILENAME, err)
	}
}

func (p *provider) close() {
	if p.client != nil {
		err := p.client.Close()
		if err != nil {
			logger.Fatal("Cannot close db file '%s' with error: %s", config.FILES_DB_APPS_FILENAME, err)
		}
	}
}
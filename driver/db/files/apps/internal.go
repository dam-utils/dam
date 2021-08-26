package apps

import (
	"os"

	"dam/driver/conf/option"
	"dam/driver/logger"
)

func (p *provider) connect() {
	var err error
	p.client, err = os.Open(option.Config.FilesDB.GetAppsFilename())
	if err != nil {
		logger.Fatal("Cannot open db file '%s' with error: %s", option.Config.FilesDB.GetAppsFilename(), err)
	}
}

func (p *provider) close() {
	if p.client != nil {
		err := p.client.Close()
		if err != nil {
			logger.Fatal("Cannot close db file '%s' with error: %s", option.Config.FilesDB.GetAppsFilename(), err)
		}
	}
}
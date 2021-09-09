package repos

import (
	"os"

	"dam/driver/conf/option"
	"dam/driver/logger"
)

func (p *provider) connect() {
	var err error
	p.client, err = os.Open(option.Config.FilesDB.GetReposFilename())
	if err != nil {
		logger.Fatal("Cannot open db file '%s' with error: %s", option.Config.FilesDB.GetReposFilename(), err)
	}
}

func (p *provider) close() {
	if p.client != nil {
		err := p.client.Close()
		if err != nil {
			logger.Fatal("Cannot close db file '%s' with error: %s", option.Config.FilesDB.GetReposFilename(), err)
		}
	}
}

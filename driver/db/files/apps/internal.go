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
package apps

import (
	"dam/config"
	"dam/driver/logger"
	"os"
)

func (p *provider) connect() {
	var err error
	p.client, err = os.Open(config.FILES_DB_APPS)
	if err != nil {
		logger.Fatal("Cannot open db file '%s' with error: %s", config.FILES_DB_APPS, err)
	}
}

func (p *provider) close() {
	if p.client != nil {
		err := p.client.Close()
		if err != nil {
			logger.Fatal("Cannot close db file '%s' with error: %s", config.FILES_DB_APPS, err)
		}
	}
}
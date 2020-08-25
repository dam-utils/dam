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
package main

import (
	"os"

	"dam/cmd"
	"dam/driver/config"
	"dam/driver/db"
	"dam/driver/engine"
)

func main() {
	// чтобы успеть закрыть все f.Close и соединения перед выходом
	defer func() {
		_ = recover()
		os.Exit(1)
	}()

	config.Prepare()
	db.Init()
	engine.Init()

	cmd.Execute()

	os.Exit(0)
}

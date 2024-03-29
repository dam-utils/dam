package main

import (
	"os"

	"dam/cmd"
	"dam/driver/conf"
	"dam/driver/db"
	"dam/driver/engine"
)

func main() {
	cmd.Init()
	// чтобы успеть закрыть все f.Close и соединения перед выходом
	defer func() {
		_ = recover()
		os.Exit(1)
	}()

	conf.Prepare()
	db.Init()
	engine.Init()

	cmd.Execute()

	os.Exit(0)
}

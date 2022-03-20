package main

import (
	"PROJECT_SCRIPT/config"
	"PROJECT_SCRIPT/handler"
	"PROJECT_SCRIPT/log"
	"PROJECT_SCRIPT/model"
)

func main() {
	if err := config.Init("", "PROJECT_SCRIPT"); err != nil {
		panic(err)
	}

	model.DB.Init()
	defer model.DB.Close()

	defer log.SyncLogger()

	handler.ClearUsers()
	// handler.Start()
}

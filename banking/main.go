package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	//log.Println("Starting out application...")
	logger.Info("Starting out application...")
	app.Start()
}

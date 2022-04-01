package main

import (
	"banking/app"
	"banking/logger"
	"fmt"
	"os"
)

func main() {
	//log.Println("Starting out application...")
	logger.Info(fmt.Sprintf("Starting server on %s:%s...", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT")))
	app.Start()
}

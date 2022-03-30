package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm_test/app"
	"log"
	"os"
	"time"
)

func main() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io writer
		logger.Config{
			SlowThreshold:             time.Second, //slow SQL threshod
			LogLevel:                  logger.Info, //log level
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: newLogger,
	})
	if err != nil {
		panic("DB 연결에 실패하였습니다.")
	}

	app.Crud(db)
	//select {}
}

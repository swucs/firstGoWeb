package main

import (
	"cron_test/app"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

func TestCron() {

	myApp := app.MyApp{}
	myApp.Print()

	c := cron.New(cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))

	i := 1
	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("Execute every minute", i)
		log.Panic("test panic")
		i++
	})

	//c.AddFunc(""+
	//	"1 * * * *", func() {
	//	fmt.Println("매분 1초마다", i)
	//	i++
	//})

	c.AddFunc("@every 30s", func() {
		fmt.Println("hello world")
	})

	c.Start()

	//time.Sleep(time.Minute * 5)
	//무한대기
	select {}
}
func main() {
	TestCron()
}

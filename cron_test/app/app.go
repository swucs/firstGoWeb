package app

import "fmt"

type MyApp struct {
}

func (a *MyApp) Print() {
	fmt.Println("My App.Print is Called.")
}

type AnotherApp struct {
}

func (a *AnotherApp) Print() {
	fmt.Println("AnotherApp.Print is Called.")
}

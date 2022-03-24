package app

import "fmt"

type MyApp struct {
}

func (a *MyApp) Print() {
	fmt.Println("My App is Called.")
}

package main

import "wire/module"

func main() {
	event := module.InitializeEvent()
	event.Start()
}

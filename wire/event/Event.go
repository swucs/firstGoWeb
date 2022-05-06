package event

import "fmt"

type Event struct {
	Greeter *Greeter
}

func NewEvent(g *Greeter) *Event {
	return &Event{Greeter: g}
}

func (e *Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

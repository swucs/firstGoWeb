package event

type Message string

type Greeter struct {
	Message Message
}

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) *Greeter {
	return &Greeter{Message: m}
}

func (g *Greeter) Greet() Message {
	return g.Message
}

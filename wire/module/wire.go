//go:build wireinject
// +build wireinject

package module

import (
	"github.com/google/wire"
	"wire/event"
)

func InitializeEvent() *event.Event {
	wire.Build(event.NewEvent, event.NewGreeter, event.NewMessage)
	return &event.Event{}
}

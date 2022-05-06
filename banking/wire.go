//go:build wireinject
// +build wireinject

package main

import (
	"banking/app"
	"banking/domain"
	"banking/service"
	"github.com/google/wire"
)

func InitializeCustomer() app.CustomerHandler {

	wire.Build(
		wire.Bind(domain.NewCustomerRepositoryDb, *app.NewClient),
		service.NewCustomerService,
		app.NewCustomerHandler)
	//
	//wire.NewSet(domain.NewCustomerRepositoryDb, service.NewCustomerService, app.NewCustomerHandler)
	//wire.Bind(new())
	return app.CustomerHandler{}
}

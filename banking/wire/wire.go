//go:build wireinject
// +build wireinject

package module

import (
	"banking/config"
	"banking/domain"
	"banking/handlers"
	"banking/service"
	"github.com/google/wire"
)

func InitializeCustomer() handlers.CustomerHandler {
	wire.Build(config.GetDbClient, domain.NewCustomerRepository, service.NewCustomerService, handlers.NewCustomerHandler)
	return handlers.CustomerHandler{}
}

func InitializeAccount() handlers.AccountHandler {
	wire.Build(config.GetDbClient, domain.NewAccountRepository, service.NewAccountService, handlers.NewAccountHandler)
	return handlers.AccountHandler{}
}

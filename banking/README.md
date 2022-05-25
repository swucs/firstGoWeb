### wire로 Dependency Inject 적용

* wire 설치
```
go get github.com/google/wire
```

* wire.go 파일 : dbClient, repository, service, handler의 인스턴스를 생성하여 서로의 의존성이 주입시키기 위한 코드이다.
```
//go:build wireinject
// +build wireinject

package module

...

func InitializeCustomer() handlers.CustomerHandler {
	wire.Build(config.GetDbClient, domain.NewCustomerRepository, service.NewCustomerService, handlers.NewCustomerHandler)
	return handlers.CustomerHandler{}
}

func InitializeAccount() handlers.AccountHandler {
	wire.Build(config.GetDbClient, domain.NewAccountRepository, service.NewAccountService, handlers.NewAccountHandler)
	return handlers.AccountHandler{}
}
```
* 위와 같이 wire.go 파일을 작성한 후 terminal에서 wire 명령어를 실행하면, 아래와 같이 wire_gen.go 파일이 생성된다.
```
// Injectors from wire.go:

func InitializeCustomer() handlers.CustomerHandler {
	db := config.GetDbClient()
	customerRepository := domain.NewCustomerRepositoryDb(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handlers.NewCustomerHandler(customerService)
	return customerHandler
}

func InitializeAccount() handlers.AccountHandler {
	db := config.GetDbClient()
	accountRepository := domain.NewAccountRepositoryDb(db)
	accountService := service.NewAccountService(accountRepository)
	accountHandler := handlers.NewAccountHandler(accountService)
	return accountHandler
}
```
* DI 하는 코드에서 다음과 같이 작성한다.
```
//wire 로 DI 주입
customerHandler := module.InitializeCustomer()
accountHandler := module.InitializeAccount()
```

* 주의사항
  * 패키지 순환 참조가 있는 경우 에러가 발생할 수 있다.

### cabin : 권한체크 RBAC 적용
@TODO 작성필요

### sqlx 적용
@TODO 작성필요
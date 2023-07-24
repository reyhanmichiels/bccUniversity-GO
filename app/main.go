package main

import (
	"bcc-university/src/business/handler/rest"
	"bcc-university/src/business/repository"
	"bcc-university/src/business/usecase"
	"bcc-university/src/sdk/conf"
	"bcc-university/src/sdk/db/sql"
)

func init() {
	conf.LoadEnv()
	sql.ConnectToDb()
	sql.Migrate()
}

func main() {
	repository := repository.InjectRepository(sql.SQLDB)
	usecase := usecase.InjectUseCase(repository)
	rest := rest.InjectRest(usecase)

	rest.Run()
}
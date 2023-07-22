package main

import (
	"bcc-university/src/sdk/conf"
	"bcc-university/src/sdk/db/sql"
)

func init() {
	conf.LoadEnv()
	sql.ConnectToDb()
}

func main() {

}
package sql

import (
	"bcc-university/src/business/entity"
)

func Migrate() {
	SQLDB.AutoMigrate(&entity.User{})
	SQLDB.AutoMigrate(&entity.Student{})
	SQLDB.AutoMigrate(&entity.Course{})
	SQLDB.AutoMigrate(&entity.Class{})
}
package sql

import (
	"bcc-university/src/business/entity"
)

func Migrate() {

	//drop if exist
	SQLDB.Migrator().DropTable("users")
	SQLDB.Migrator().DropTable("students")
	SQLDB.Migrator().DropTable("courses")
	SQLDB.Migrator().DropTable("classes")
	SQLDB.Migrator().DropTable("user_classes")

	//migrate
	SQLDB.AutoMigrate(&entity.User{})
	SQLDB.AutoMigrate(&entity.Student{})
	SQLDB.AutoMigrate(&entity.Course{})
	SQLDB.AutoMigrate(&entity.Class{})

}
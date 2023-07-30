package sql

import (
	"bcc-university/src/business/entity"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {

	//course
	course := entity.Course{
		Name:   "Pemrograman Web Lanjut",
		Credit: 4,
	}
	SQLDB.Create(&course)

	course = entity.Course{
		Name:   "Pemrograman Dasar",
		Credit: 4,
	}
	SQLDB.Create(&course)

	course = entity.Course{
		Name:   "Pemrograman Lanjut",
		Credit: 4,
	}
	SQLDB.Create(&course)

	//classes
	class := entity.Class{
		Name:      "Pemrograman Web Lanjut A",
		Course_id: 1,
	}
	SQLDB.Create(&class)

	class = entity.Class{
		Name:      "Pemrograman Web Lanjut B",
		Course_id: 1,
	}
	SQLDB.Create(&class)

	class = entity.Class{
		Name:      "Pemrograman Dasar A",
		Course_id: 2,
	}
	SQLDB.Create(&class)

	class = entity.Class{
		Name:      "Pemrograman Dasar B",
		Course_id: 2,
	}
	SQLDB.Create(&class)
	class = entity.Class{
		Name:      "Pemrograman Lanjut A",
		Course_id: 3,
	}
	SQLDB.Create(&class)

	class = entity.Class{
		Name:      "Pemrograman Lanjut B",
		Course_id: 3,
	}
	SQLDB.Create(&class)

	//user
	pass, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)

	user := entity.User{
		Name:     "admin",
		Username: "admin",
		Role:     "admin",
		Email:    "admin@gmail.com",
		Password: string(pass),
	}
	SQLDB.Create(&user)

	user = entity.User{
		Name:     "Reyhan Hafiz Rusyard",
		Username: "reyhan",
		Email:    "reyhan@gmail.com",
		Password: string(pass),
	}
	SQLDB.Create(&user)

	user = entity.User{
		Name:     "Benjamin Franklin",
		Username: "Ben",
		Email:    "benjamin@gmail.com",
		Password: string(pass),
	}
	SQLDB.Create(&user)

	user = entity.User{
		Name:     "Lebron James",
		Username: "Lebron",
		Email:    "lebron@gmail.com",
		Password: string(pass),
	}
	SQLDB.Create(&user)

	//student
	student := entity.Student{
		User_id:           2,
		Student_id_number: "210101",
	}
	SQLDB.Create(&student)

	student = entity.Student{
		User_id:           3,
		Student_id_number: "210102",
	}
	SQLDB.Create(&student)

	student = entity.Student{
		User_id:           4,
		Student_id_number: "210103",
	}
	SQLDB.Create(&student)

	userClass := entity.UserClass{
		UserID:  2,
		ClassID: 1,
	}
	SQLDB.Create(&userClass)

	userClass = entity.UserClass{
		UserID:  2,
		ClassID: 3,
	}
	SQLDB.Create(&userClass)

	userClass = entity.UserClass{
		UserID:  2,
		ClassID: 5,
	}
	SQLDB.Create(&userClass)

	userClass = entity.UserClass{
		UserID:  3,
		ClassID: 2,
	}
	SQLDB.Create(&userClass)

	userClass = entity.UserClass{
		UserID:  3,
		ClassID: 4,
	}
	SQLDB.Create(&userClass)

	userClass = entity.UserClass{
		UserID:  3,
		ClassID: 6,
	}
	SQLDB.Create(&userClass)

	userClass = entity.UserClass{
		UserID:  4,
		ClassID: 1,
	}
	SQLDB.Create(&userClass)

	userClass = entity.UserClass{
		UserID:  4,
		ClassID: 4,
	}
	SQLDB.Create(&userClass)

	userClass = entity.UserClass{
		UserID:  4,
		ClassID: 5,
	}
	SQLDB.Create(&userClass)

}

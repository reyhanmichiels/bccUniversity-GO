package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/repository"
	"bcc-university/src/sdk/library"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type StudentUseCase interface {
	ClaimStudentNumber(loginUser entity.User) (entity.StudentResponse, interface{})
}

type studentUseCase struct {
	studentRepository repository.StudentRepository
}

func NewStudentUseCase(studentRepository repository.StudentRepository) StudentUseCase {
	return &studentUseCase{
		studentRepository: studentRepository,
	}
}

func (studentUseCase *studentUseCase) ClaimStudentNumber(loginUser entity.User) (entity.StudentResponse, interface{}) {

	//validate if user has student number
	if len(loginUser.Student.Student_id_number) != 0 {

		errorObject := library.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "you already have student number",
			Err:     errors.New("this endpoint only can called for student who doesnt have student number"),
		}
		return entity.StudentResponse{}, errorObject

	}

	//get last student number
	lastStudentNumber, err := studentUseCase.studentRepository.GetLastStudentNumber()
	if err != nil {

		errorObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to get last student number",
			Err:     err,
		}
		return entity.StudentResponse{}, errorObject

	}

	//generate new student number
	var newStudentNumber string
	var idString string
	var batchString string

	yearString := strconv.Itoa(time.Now().Year())[:2]
	newStudentNumber += yearString

	idNumber, err := strconv.Atoi(lastStudentNumber[4:])
	if err != nil {

		errorObject := library.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "failed covert string to number",
			Err:     err,
		}
		return entity.StudentResponse{}, errorObject

	}

	if idNumber == 99 {

		idString = "01"

		batchNumber, err := strconv.Atoi(lastStudentNumber[2:4])
		if err != nil {

			errorObject := library.ErrorObject{
				Code:    http.StatusBadRequest,
				Message: "failed covert string to number",
				Err:     err,
			}
			return entity.StudentResponse{}, errorObject

		}

		batchNumber += 1
		if batchNumber < 10 {

			batchString = "0" + strconv.Itoa(batchNumber)

		} else {

			batchString = strconv.Itoa(batchNumber)

		}

	} else {

		batchString = lastStudentNumber[2:4]

		idNumber += 1
		if idNumber < 10 {

			idString = "0" + strconv.Itoa(idNumber)

		} else {

			idString = strconv.Itoa(idNumber)

		}

	}

	newStudentNumber += batchString + idString

	//insert user to student
	student := entity.Student{
		User_id:           loginUser.ID,
		Student_id_number: newStudentNumber,
	}

	err = studentUseCase.studentRepository.CreateStudent(student)
	if err != nil {

		errorObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to make new Student",
			Err:     err,
		}
		return entity.StudentResponse{}, errorObject

	}

	return entity.StudentResponse{Student_id_number: newStudentNumber}, nil

}

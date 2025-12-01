package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

type StudentReponse struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Name      string    `json:"name"`
	CPF       int       `json:"cpf"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Active    bool      `json:"active"`
}

func NewResponse(students []Student) []StudentReponse {
	studentsReponse := []StudentReponse{}

	for _, student := range students {
		studentReponse := StudentReponse{

			ID:        int(student.ID),
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			Name:      student.Name,
			Email:     student.Email,
			CPF:       student.CPF,
			Age:       student.Age,
			Active:    student.Active,
		}

		studentsReponse = append(studentsReponse, studentReponse)
	}

	return studentsReponse

}

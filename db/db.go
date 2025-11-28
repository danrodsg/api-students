package db

import (
	"fmt"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"registration"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)

	}

	db.AutoMigrate(&Student{})

	return db

}

func AddStudent(student Student) error {
	db := Init()

	result := db.Create(&student)
	if result.Error != nil {
		return result.Error

	}

	fmt.Println("Create student")
	return nil

}

func GetStudents() ([]Student, error) {
	students := []Student{}

	db := Init()
	err := db.Find(&students).Error
	return students, err
}

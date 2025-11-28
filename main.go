package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"github.com/danrodsg/api-students/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/students", getStudent)
  e.POST("/students", createStudent)
  e.GET("/students/:id",getStudent)
  e.PUT("/students/:id", updateStudent)
  e.DELETE("/students/:id", deleteStudent)

  // Start server
  if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
    slog.Error("failed to start server", "error", err)
  }
}

// Handler

func getStudents(c echo.Context) error {
  return c.String(http.StatusOK, "List of all students")
}

func createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
	    return err
	}
	db.AddStudent(student)
  return c.String(http.StatusOK, "Create student")
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	getStud :=fmt.Sprintf("Get %s student", id)
  return c.String(http.StatusOK, getStud)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	UpdateStud :=fmt.Sprintf("Update %s student", id)
    return c.String(http.StatusOK, UpdateStud)
	
}

func deleteStudent(c echo.Context) error {

    id := c.Param("id")
	deleteStud :=fmt.Sprintf("delete %s student", id)
    return c.String(http.StatusOK, deleteStud)
}

package api

import (
	"github.com/danrodsg/api-students/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/swaggo/echo-swagger"
	_"github.com/danrodsg/api-students/docs"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

//@title Student API
//@version 1.0
//@descrpition This is a sample server Student API
//@host localhost:8080
//@BasePath / 
//@schemes http

func NewServer() *API {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentsDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentsDB,
	}
}

func (api *API) Start() error {

	return api.Echo.Start(":8080")

}

func (api *API) ConfigureRoutes() {

	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
	api.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

}

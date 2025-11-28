package api

import(
	"fmt"
	"net/http"
	"github.com/danrodsg/api-students/db"

	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	
	
)

type API struct {

	Echo *echo.Echo
	DB *db.StudentHandler
}

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


func(api *API) ConfigureRoutes(){

	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)

}


func (api *API)getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get student")

	}
	return c.JSON(http.StatusOK, students)

}

func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}

	return c.String(http.StatusOK, "Create student")
}

func (api *API) getStudent(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("Get %s student", id)
	return c.String(http.StatusOK, getStud)
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	UpdateStud := fmt.Sprintf("Update %s student", id)
	return c.String(http.StatusOK, UpdateStud)

}

func (api *API) deleteStudent(c echo.Context) error {

	id := c.Param("id")
	deleteStud := fmt.Sprintf("delete %s student", id)
	return c.String(http.StatusOK, deleteStud)
}

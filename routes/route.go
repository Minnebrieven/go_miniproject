package routes

import (
	"swim-class/constants"
	"swim-class/controllers"
	m "swim-class/middlewares"
	"swim-class/repositories"
	"swim-class/services"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(db *gorm.DB) *echo.Echo {
	//echo instance
	e := echo.New()

	//REPOSITORIES
	userRepository := repositories.NewUserRepository(db)
	instructorRepository := repositories.NewInstructorRepository(db)

	//SERVICES
	userService := services.NewUserService(userRepository)
	instructorService := services.NewInstructorService(instructorRepository)

	//CONTROLLERS
	userController := controllers.NewUserController(userService)
	instructorController := controllers.NewInstructorController(instructorService)

	//Log
	m.LogMiddleware(e)

	//ROUTES

	//------------------NOT AUTHENTICATED--------------------------------
	//USERS ROUTES
	e.POST("/login", userController.Login)
	e.POST("/register", userController.CreateUser)

	//INSTRUCTORS ROUTES
	e.GET("/instructors/all", instructorController.GetAllInstructors)
	e.GET("/instructors/:id", instructorController.GetInstructorByID)

	//------------------AUTHENTICATED------------------------------------
	//USERS ROUTES
	usersJWT := e.Group("/users")
	usersJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	usersJWT.GET("", userController.GetAllUsers)
	usersJWT.GET("/:id", userController.GetUserByID)
	usersJWT.PUT("/:id", userController.EditUser)
	usersJWT.DELETE("/:id", userController.DeleteUser)

	//INSTRUCTORS ROUTES
	instructorsJWT := e.Group("/instructors")
	instructorsJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	instructorsJWT.POST("", instructorController.CreateInstructor)
	instructorsJWT.PUT("/:id", instructorController.EditInstrutor)
	instructorsJWT.DELETE("/:id", instructorController.DeleteInstructor)

	return e
}

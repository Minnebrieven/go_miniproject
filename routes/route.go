package routes

import (
	"os"
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
	classRepository := repositories.NewClassRepository(db)
	classCategoryRepository := repositories.NewClassCategoryRepository(db)
	classParticipantRepository := repositories.NewClassParticipantRepository(db)

	//SERVICES
	userService := services.NewUserService(userRepository)
	instructorService := services.NewInstructorService(instructorRepository)
	classService := services.NewClassService(classRepository, classCategoryRepository, instructorRepository)
	classCategoryService := services.NewClassCategoryService(classCategoryRepository)
	classParticipantService := services.NewClassParticipantService(classParticipantRepository, classRepository, userRepository)

	//CONTROLLERS
	userController := controllers.NewUserController(userService)
	instructorController := controllers.NewInstructorController(instructorService)
	classController := controllers.NewClassController(classService)
	classCategoryController := controllers.NewClassCategoryController(classCategoryService)
	classParticipantController := controllers.NewClassParticipantController(classParticipantService)

	//Log
	m.LogMiddleware(e)

	//=======================ROUTES=====================================

	//------------------NOT AUTHENTICATED--------------------------------
	//USERS ROUTES
	e.POST("/login", userController.Login)
	e.POST("/register", userController.CreateUser)

	//INSTRUCTORS ROUTES
	e.GET("/instructors/all", instructorController.GetAllInstructors)
	e.GET("/instructors/:id", instructorController.GetInstructorByID)

	// CLASS CATEGORY ROUTES
	e.GET("/category/all", classCategoryController.GetAllClassCategories)
	e.GET("/category/:id", classCategoryController.GetClassCategoryByID)

	//CLASS ROUTES
	e.GET("/classes/all", classController.GetAllClasses)
	e.GET("/classes/:id", classController.GetClassByID)

	//------------------AUTHENTICATED------------------------------------
	//USERS ROUTES
	usersJWT := e.Group("/users")
	usersJWT.Use(mid.JWT([]byte(os.Getenv("JWT_KEY"))))
	usersJWT.GET("", userController.GetAllUsers, m.IsAdmin)
	usersJWT.GET("/:id", userController.GetUserByID)
	usersJWT.PUT("/:id", userController.EditUser)
	usersJWT.DELETE("/:id", userController.DeleteUser)

	//INSTRUCTORS ROUTES
	instructorsJWT := e.Group("/instructors")
	instructorsJWT.Use(mid.JWT([]byte(os.Getenv("JWT_KEY"))), m.IsAdmin)
	instructorsJWT.POST("", instructorController.CreateInstructor)
	instructorsJWT.PUT("/:id", instructorController.EditInstrutor)
	instructorsJWT.DELETE("/:id", instructorController.DeleteInstructor)

	//CLASSES ROUTES
	classesJWT := e.Group("/classes")
	classesJWT.Use(mid.JWT([]byte(os.Getenv("JWT_KEY"))), m.IsAdmin)
	classesJWT.POST("", classController.CreateClass)
	classesJWT.PUT("/:id", classController.EditClass)
	classesJWT.DELETE("/:id", classController.DeleteClass)

	//CLASS CATEGORY ROUTES
	categoryJWT := e.Group("/category")
	categoryJWT.Use(mid.JWT([]byte(os.Getenv("JWT_KEY"))), m.IsAdmin)
	categoryJWT.POST("", classCategoryController.CreateClassCategory)
	categoryJWT.PUT("/:id", classCategoryController.EditClassCategory)
	categoryJWT.DELETE("/:id", classCategoryController.DeleteClassCategory)

	// CLASS PARTICIPANT ROUTES
	participantJWT := e.Group("/participants")
	participantJWT.Use(mid.JWT([]byte(os.Getenv("JWT_KEY"))))
	participantJWT.GET("/all", classParticipantController.GetAllClassParticipants)
	participantJWT.GET("/myclass/:user", classParticipantController.GetAllClassParticipantsByUserID)
	participantJWT.GET("/:id", classParticipantController.GetClassParticipantByID)
	participantJWT.POST("", classParticipantController.CreateClassParticipant)
	participantJWT.PUT("/:id", classParticipantController.EditClassParticipant)
	participantJWT.DELETE("/:id", classParticipantController.DeleteClassParticipant)

	return e
}

package controllers

import (
	"net/http"
	"strconv"
	"swim-class/models"
	"swim-class/services"

	"github.com/labstack/echo/v4"
)

type (
	UserController interface{}

	userController struct {
		userService services.UserService
	}
)

func NewUserController(userServ services.UserService) *userController {
	return &userController{userService: userServ}
}

func (u *userController) GetAllUsers(c echo.Context) error {
	users, err := u.userService.GetAllUsersService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   users,
	})
}

func (u *userController) GetUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	var user = models.User{ID: uint(userID)}
	user, err = u.userService.GetUserService(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   user,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	if err := u.userService.CreateUserService(user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   user,
	})
}

func (u *userController) EditUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	modifiedUserData := models.User{}
	if err := c.Bind(&modifiedUserData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	user, err := u.userService.EditUserService(userID, modifiedUserData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   user,
	})

}

func (u *userController) DeleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	if err = u.userService.DeleteUserService(userID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"id":     userID,
	})
}

func (u *userController) Login(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	token, err := u.userService.Login(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"token":  token,
	})
}

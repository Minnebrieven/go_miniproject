package controllers

import (
	"net/http"
	"strconv"
	"swim-class/dto"
	"swim-class/middlewares"
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
			"error": err.Error(),
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

	var user = dto.UserDTO{ID: userID}
	user, err = u.userService.GetUserService(user)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusOK, echo.Map{
				"error": err.Error(),
			})
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   user,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	userDTO := dto.UserDTO{}

	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	userDTO, err := u.userService.CreateUserService(userDTO)
	if err != nil {
		if err.Error() == "duplicate key found" {
			return c.JSON(http.StatusConflict, echo.Map{
				"error": err.Error(),
			})
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
	}

	token, err := middlewares.CreateToken(userDTO.ID, userDTO.Email, false)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   userDTO,
		"token":  token,
	})
}

func (u *userController) EditUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	isSameUser := middlewares.IsSameUser(c, float64(userID))
	if !isSameUser {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized - not the same user to edit",
		})
	}

	modifiedUserData := dto.UserDTO{}
	if err := c.Bind(&modifiedUserData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	user, err := u.userService.EditUserService(userID, modifiedUserData)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusOK, echo.Map{
				"error": err.Error(),
			})
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
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

	err = u.userService.DeleteUserService(userID)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusOK, echo.Map{
				"error": err.Error(),
			})
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"id":     userID,
	})
}

func (u *userController) Login(c echo.Context) error {
	user := dto.UserDTO{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	user, token, err := u.userService.Login(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   user,
		"token":  token,
	})
}

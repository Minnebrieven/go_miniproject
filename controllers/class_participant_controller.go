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
	ClassParticipantController interface{}

	classParticipantController struct {
		classParticipantService services.ClassParticipantService
	}
)

func NewClassParticipantController(classParticipantServ services.ClassParticipantService) *classParticipantController {
	return &classParticipantController{classParticipantService: classParticipantServ}
}

func (cpc *classParticipantController) GetAllClassParticipants(c echo.Context) error {
	classParticipants, err := cpc.classParticipantService.GetAllClassParticipantsService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   classParticipants,
	})
}

func (cpc *classParticipantController) GetAllClassParticipantsByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	isSameUser := middlewares.IsSameUser(c, float64(userID))
	if !isSameUser {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized - not the same user",
		})
	}

	classParticipants, err := cpc.classParticipantService.GetAllClassParticipantsByUserIDService(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   classParticipants,
	})
}

func (cpc *classParticipantController) GetClassParticipantByID(c echo.Context) error {
	classParticipantID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	var classParticipant = dto.ClassParticipantDTO{ID: classParticipantID}
	classParticipant, err = cpc.classParticipantService.GetClassParticipantService(classParticipant)
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
		"data":   classParticipant,
	})
}

func (cpc *classParticipantController) CreateClassParticipant(c echo.Context) error {
	classParticipantDTO := dto.ClassParticipantDTO{}

	if err := c.Bind(&classParticipantDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	userID := classParticipantDTO.User.ID
	isSameUser := middlewares.IsSameUser(c, float64(userID))
	if !isSameUser {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized - not the same user",
		})
	}

	classParticipantDTO, err := cpc.classParticipantService.CreateClassParticipantService(classParticipantDTO)
	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusBadRequest, echo.Map{
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
		"data":   classParticipantDTO,
	})
}

func (cpc *classParticipantController) EditClassParticipant(c echo.Context) error {
	classParticipantID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	isSameUser := middlewares.IsSameUser(c, float64(classParticipantID))
	if !isSameUser {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized - not the same user",
		})
	}

	modifiedClassParticipantData := dto.ClassParticipantDTO{}
	if err := c.Bind(&modifiedClassParticipantData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	classParticipant, err := cpc.classParticipantService.EditClassParticipantService(classParticipantID, modifiedClassParticipantData)
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
		"data":   classParticipant,
	})

}

func (cpc *classParticipantController) DeleteClassParticipant(c echo.Context) error {
	classParticipantID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	isSameUser := middlewares.IsSameUser(c, float64(classParticipantID))
	if !isSameUser {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized - not the same user",
		})
	}

	err = cpc.classParticipantService.DeleteClassParticipantService(classParticipantID)
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
		"id":     classParticipantID,
	})
}

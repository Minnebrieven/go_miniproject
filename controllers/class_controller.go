package controllers

import (
	"net/http"
	"strconv"
	"swim-class/dto"
	"swim-class/services"

	"github.com/labstack/echo/v4"
)

type (
	ClassController interface{}

	classController struct {
		classService services.ClassService
	}
)

func NewClassController(classServ services.ClassService) *classController {
	return &classController{classService: classServ}
}

func (cl *classController) GetAvailableClasses(c echo.Context) error {
	classes, err := cl.classService.GetAvailableClasses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   classes,
	})
}

func (cl *classController) GetAllClasses(c echo.Context) error {
	classes, err := cl.classService.GetAllClassesService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   classes,
	})
}

func (cl *classController) GetClassByID(c echo.Context) error {
	classID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	var class = dto.ClassDTO{ID: classID}
	class, err = cl.classService.GetClassService(class)
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
		"data":   class,
	})
}

func (cl *classController) CreateClass(c echo.Context) error {
	classDTO := dto.ClassDTO{}

	if err := c.Bind(&classDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	classDTO, err := cl.classService.CreateClassService(classDTO)
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
		"data":   classDTO,
	})
}

func (cl *classController) EditClass(c echo.Context) error {
	classID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	modifiedClassData := dto.ClassDTO{}
	if err := c.Bind(&modifiedClassData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	class, err := cl.classService.EditClassService(classID, modifiedClassData)
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
		"data":   class,
	})

}

func (cl *classController) DeleteClass(c echo.Context) error {
	classID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	err = cl.classService.DeleteClassService(classID)
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
		"id":     classID,
	})
}

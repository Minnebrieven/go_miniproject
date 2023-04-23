package controllers

import (
	"net/http"
	"strconv"
	"swim-class/models"
	"swim-class/services"

	"github.com/labstack/echo/v4"
)

type (
	InstructorController interface{}

	instructorController struct {
		instructorService services.InstructorService
	}
)

func NewInstructorController(instructorServ services.InstructorService) *instructorController {
	return &instructorController{instructorService: instructorServ}
}

func (i *instructorController) GetAllInstructors(c echo.Context) error {
	instructors, err := i.instructorService.GetAllInstructorsService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   instructors,
	})
}

func (i *instructorController) GetInstructorByID(c echo.Context) error {
	instructorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	var instructor = models.Instructor{ID: uint(instructorID)}
	instructor, err = i.instructorService.GetInstructorService(instructor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   instructor,
	})
}

func (i *instructorController) CreateInstructor(c echo.Context) error {
	instructor := models.Instructor{}
	if err := c.Bind(&instructor); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	if err := i.instructorService.CreateInstructorService(instructor); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   instructor,
	})
}

func (i *instructorController) EditInstrutor(c echo.Context) error {
	instructorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	modifiedInstructorData := models.Instructor{}
	if err := c.Bind(&modifiedInstructorData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	instructor, err := i.instructorService.EditInstructorService(instructorID, modifiedInstructorData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   instructor,
	})

}

func (i *instructorController) DeleteInstructor(c echo.Context) error {
	instructorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	if err = i.instructorService.DeleteInstructorService(instructorID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"id":     instructorID,
	})
}

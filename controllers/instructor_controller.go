package controllers

import (
	"net/http"
	"strconv"
	"swim-class/dto"
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
			"error": err.Error(),
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

	var instructor = dto.InstructorDTO{ID: instructorID}
	instructor, err = i.instructorService.GetInstructorService(instructor)
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
		"data":   instructor,
	})
}

func (i *instructorController) CreateInstructor(c echo.Context) error {
	instructor := dto.InstructorDTO{}
	var err error
	if err = c.Bind(&instructor); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	if instructor, err = i.instructorService.CreateInstructorService(instructor); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
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

	modifiedInstructorData := dto.InstructorDTO{}
	if err := c.Bind(&modifiedInstructorData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	instructor, err := i.instructorService.EditInstructorService(instructorID, modifiedInstructorData)
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

	err = i.instructorService.DeleteInstructorService(instructorID)
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
		"id":     instructorID,
	})
}

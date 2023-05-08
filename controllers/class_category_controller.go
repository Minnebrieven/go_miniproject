package controllers

import (
	"net/http"
	"strconv"
	"swim-class/dto"
	"swim-class/services"

	"github.com/labstack/echo/v4"
)

type (
	ClassCategoryController interface{}

	classCategoryController struct {
		classCategoryService services.ClassCategoryService
	}
)

func NewClassCategoryController(classCategoryServ services.ClassCategoryService) *classCategoryController {
	return &classCategoryController{classCategoryService: classCategoryServ}
}

func (ccc *classCategoryController) GetAllClassCategories(c echo.Context) error {
	classCategories, err := ccc.classCategoryService.GetAllClassCategoriesService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   classCategories,
	})
}

func (ccc *classCategoryController) GetClassCategoryByID(c echo.Context) error {
	classCategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	var classCategory = dto.ClassCategoryDTO{ID: classCategoryID}
	classCategory, err = ccc.classCategoryService.GetClassCategoryService(classCategory)
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
		"data":   classCategory,
	})
}

func (ccc *classCategoryController) CreateClassCategory(c echo.Context) error {
	classCategoryDTO := dto.ClassCategoryDTO{}

	if err := c.Bind(&classCategoryDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	classCategoryDTO, err := ccc.classCategoryService.CreateClassCategoryService(classCategoryDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   classCategoryDTO,
	})
}

func (ccc *classCategoryController) EditClassCategory(c echo.Context) error {
	classCategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	modifiedClassCategoryData := dto.ClassCategoryDTO{}
	if err := c.Bind(&modifiedClassCategoryData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	classCategory, err := ccc.classCategoryService.EditClassCategoryService(classCategoryID, modifiedClassCategoryData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   classCategory,
	})

}

func (ccc *classCategoryController) DeleteClassCategory(c echo.Context) error {
	classCategoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	if err = ccc.classCategoryService.DeleteClassCategoryService(classCategoryID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"id":     classCategoryID,
	})
}

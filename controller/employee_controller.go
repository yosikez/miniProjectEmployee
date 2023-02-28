package controller

import (
	"miniProject/database"
	"miniProject/helper/validation"
	"miniProject/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct{}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{}
}

func (em *EmployeeController) FindAll(c *gin.Context) {
	var employees []model.Employee

	if err := database.DB.Find(&employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to find employees",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   employees,
	})
}

func (em *EmployeeController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid employee id",
		})
		return
	}

	var employee model.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "employee not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : employee,
	})
}

func (em *EmployeeController) Create(c *gin.Context) {
	var employee model.Employee

	if err := c.ShouldBind(&employee); err != nil {

		errFields := validation.GetErrMess(err)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": errFields,
		})
		return

	}

	if err := database.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data" : employee,
	})
}

func (em *EmployeeController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid employee id",
		})
		return
	}

	var employee model.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "employee not found",
		})
		return
	}

	if err := c.ShouldBind(&employee); err != nil {
		errFields := validation.GetErrMess(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": errFields,
		})
		return
	}
	
	if err := database.DB.Save(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update employee",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   employee,
	})
}

func (em *EmployeeController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid employee id",
		})
		return
	}

	var employee model.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "employee not found",
		})
		return
	}

	if err := database.DB.Delete(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete employee",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "employee deleted successfully",
	})
}

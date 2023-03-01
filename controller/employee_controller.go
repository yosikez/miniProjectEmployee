package controller

import (
	"fmt"
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

	result := database.DB.Find(&employees)

	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to find all employee",
			"error": err.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to find all employee",
			"error": "record not found",
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
			"error": "id must be a number",
		})
		return
	}

	var employee model.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to find employee",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : employee,
	})
}

func (em *EmployeeController) Create(c *gin.Context) {
	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		errFields := validation.GetErrMess(err)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "validation error",
			"error": errFields,
		})
		return

	}

	if err := database.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create user",
			"error" : err.Error(),
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
			"error": "id must be a number",
		})
		return
	}

	var employee model.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to find employee to update",
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBind(&employee); err != nil {
		errFields := validation.GetErrMess(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "validation error",
			"error": errFields,
		})
		return
	}
	
	if err := database.DB.Save(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update employee",
			"error": err.Error(),
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
			"error": "id must be a number",
		})
		fmt.Println(err)
		return
	}

	var employee model.Employee
	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to find employee to delete",
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Delete(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete employee",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "employee deleted successfully",
	})
}

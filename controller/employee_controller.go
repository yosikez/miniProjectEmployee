package controller

import (
	"miniProject/database"
	"miniProject/helper/validation"
	"miniProject/inputs"
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
	var employeeInput inputs.EmployeeInput

	if err := c.ShouldBind(&employeeInput); err != nil {

		errFields := validation.GetErrMess(err)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": errFields,
		})
		return

	}

	employee := model.Employee{
		Name:    employeeInput.Name,
		Email:   employeeInput.Email,
		Address: employeeInput.Address,
		Phone:   employeeInput.Phone,
		Gender:  model.Gender(employeeInput.Gender),
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

	var employeeInput inputs.EmployeeInput
	if err := c.ShouldBind(&employeeInput); err != nil {
		errFields := validation.GetErrMess(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": errFields,
		})
		return
	}

	employee.Name = employeeInput.Name
	employee.Email = employeeInput.Email
	employee.Address = employeeInput.Address
	employee.Phone = employeeInput.Phone
	employee.Gender = model.Gender(employeeInput.Gender)

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

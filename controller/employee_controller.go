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
	var Employees []model.Employee

	if err := database.DB.Find(&Employees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   Employees,
	})
}

func (em *EmployeeController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee id",
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
		"status": "OK",
		"data":   employee,
	})
}

func (em *EmployeeController) Create(c *gin.Context) {
	var employeeInput inputs.EmployeeInput

	if err := c.ShouldBind(&employeeInput); err != nil {

		errFields := validation.GetErrMess(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errFields,
		})
		return

	}

	employee := model.Employee{
		Name: employeeInput.Name,
		Email: employeeInput.Email,
		Address: employeeInput.Address,
		Phone: employeeInput.Phone,
		Gender: model.Gender(employeeInput.Gender),
	}

	if err := database.DB.Create(&employee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": "Created",
		"data":    employee,
	})
}

func (em *EmployeeController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee id",
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
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errFields,
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
			"error": "failed to update employee",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Updated",
		"data":   employee,
	})
}

func (em *EmployeeController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee id",
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
			"error": "failed to delete employee",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "employee deleted successfully",
	})
}

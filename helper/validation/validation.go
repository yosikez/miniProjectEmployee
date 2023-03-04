package validation

import (
	"errors"
	"miniProject/database"
	"miniProject/model"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("gender", GenderValidator)
		v.RegisterValidation("uniqueMail", UniqueField)
	}
}

func UniqueField(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	var employee model.Employee
	result := database.DB.Table("employees").Where("email = ?", email).First(&employee)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}

	parentIDField := fl.Parent().FieldByName("ID")
	if !parentIDField.IsValid() {
		return false
	}

	parentID := parentIDField.Interface().(int64)

	duplicateEmployee := model.Employee{}
	duplicateResult := database.DB.Table("employees").Where("email = ?", email).First(&duplicateEmployee)
	if !errors.Is(duplicateResult.Error, gorm.ErrRecordNotFound) && duplicateEmployee.ID != int64(parentID) {
		return false
	}

	return true
}

func GenderValidator(fl validator.FieldLevel) bool {
	gender := fl.Field().String()
	return gender == "male" || gender == "female"
}
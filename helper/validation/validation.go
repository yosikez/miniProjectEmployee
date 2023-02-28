package validation

import (
	"errors"
	"fmt"
	"miniProject/database"
	"miniProject/model"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterCustomValidator(){
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

func GetErrMess(err error) map[string]string {
	errFields := make(map[string]string)
	var errs validator.ValidationErrors

	if errors.As(err, &errs) {
		for _, errField := range errs {
			switch errField.Tag() {
			case "required":
				errFields[errField.Field()] = fmt.Sprintf("%s is required", errField.Field())
			case "email":
				errFields[errField.Field()] = fmt.Sprintf("%s is not a valid email address", errField.Field())
			case "min":
				errFields[errField.Field()] = fmt.Sprintf("%s must be at least %s characters/nums long", errField.Field(), errField.Param())
			case "max":
				errFields[errField.Field()] = fmt.Sprintf("%s must be at most %s characters/nums long", errField.Field(), errField.Param())
			case "gender":
				errFields[errField.Field()] = fmt.Sprintf("%s must be male or female", errField.Tag())
			case "uniqueMail":
				errFields[errField.Field()] = fmt.Sprintf("%s already taken", errField.Field())
			default:
				errFields[errField.Field()] = fmt.Sprintf("%s is not valid", errField.Field())
			}
		}
	}

	return errFields
}

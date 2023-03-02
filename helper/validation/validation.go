package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"miniProject/database"
	"miniProject/helper"
	"miniProject/input"
	"miniProject/model"
	"reflect"

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

func GetErrMess(err error) map[string]string {
	errFields := make(map[string]string)
	var errs validator.ValidationErrors

	if errors.As(err, &errs) {
		for _, errField := range errs {
			var field, disField string

			mapStr, isTrue := helper.CheckDiveField(errField)

			if isTrue {
				field = fmt.Sprintf("%s.%s.%s", mapStr["field"], mapStr["index"], mapStr["attribute"])
				disField = mapStr["attribute"]
			} else {
				structField, _ :=  reflect.TypeOf(input.JsonDataOpportunity{}).FieldByName(errField.Field())
				field = helper.GetJSONTagName(structField)
				disField = field
			}


			switch errField.Tag() {
			case "required":
				errFields[field] = fmt.Sprintf("%s is required", disField)
			case "email":
				errFields[field] = fmt.Sprintf("%s is not a valid email address", disField)
			case "min":
				errFields[field] = fmt.Sprintf("%s must be at least %s characters/nums long", disField, errField.Param())
			case "max":
				errFields[field] = fmt.Sprintf("%s must be at most %s characters/nums long", disField, errField.Param())
			case "gender":
				errFields[field] = fmt.Sprintf("%s must be male or female", disField)
			case "uniqueMail":
				errFields[field] = fmt.Sprintf("%s already taken", disField)
			case "numeric":
				errFields[field] = fmt.Sprintf("%s must be numeric", disField)
			case "datetime":
				errFields[field] = fmt.Sprintf("%s format must be yyyy-mm-dd hh:mm:ss", disField)
			default:
				errFields[field] = fmt.Sprintf("%s is not valid", disField)
			}
		}
	} else {
		var unmarshalErr *json.UnmarshalTypeError
		if errors.As(err, &unmarshalErr) {
			errFields[unmarshalErr.Field] = fmt.Sprintf("Invalid type. Expected %v but got %v", unmarshalErr.Type, unmarshalErr.Value)
		} else {
			errFields["input"] = err.Error()
		}
	}

	return errFields
}
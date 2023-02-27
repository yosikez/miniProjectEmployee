package validation

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidator(){
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("gender", GenderValidator)
	}
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
				errFields[errField.Field()] = fmt.Sprintf("%s must be at least %s characters long", errField.Field(), errField.Param())
			case "max":
				errFields[errField.Field()] = fmt.Sprintf("%s must be at most %s characters long", errField.Field(), errField.Param())
			case "gender":
				errFields[errField.Field()] = fmt.Sprintf("%s must be male or female", errField.Tag())
			default:
				errFields[errField.Field()] = fmt.Sprintf("%s is not valid", errField.Field())
			}
		}
	}

	return errFields
}

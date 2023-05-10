package utils

import (
	"github.com/go-playground/validator"
)

func ValidateRequest(validate *validator.Validate, request interface{}) (error, []string) {
	var errorStrings []string
	err := validate.Struct(request)
	var errorString string
	if err != nil {
		for _, errorValidation := range err.(validator.ValidationErrors) {
			errorString = errorValidation.Field() + " is " + errorValidation.Tag()
			errorStrings = append(errorStrings, errorString)
		}
		return err, errorStrings
	}

	return nil, []string{}
}

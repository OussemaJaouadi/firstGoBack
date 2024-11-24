package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ParseValidationErrors(err error) []string {
	var validationErrors []string

	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range fieldErrors {
			validationErrors = append(validationErrors, fmt.Sprintf("Field '%s' failed validation", fe.Field()))
		}
	} else {
		validationErrors = append(validationErrors, err.Error())
	}

	return validationErrors
}

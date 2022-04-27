package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	fmt.Println("error")
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e.Error())
		errors = append(errors, e.Error())
	}
	return errors
}
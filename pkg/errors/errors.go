package errors

import (
	"errors"
	"gin-blog/app/providers/validation"
	"github.com/go-playground/validator/v10"
	"strings"
)

var errorsList = make(map[string]string)

var errorMessages = validation.ErrorMessages()

func GetErrorMsg(fe validator.FieldError) string {
	return errorMessages[fe.Tag()]
}

func Init() {
	errorsList = map[string]string{}
}

func AddFromErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), GetErrorMsg(fieldError))
		}
	}
}

func Add(key string, value string) {
	errorsList[strings.ToLower(key)] = value
}

func Get() map[string]string {
	return errorsList
}

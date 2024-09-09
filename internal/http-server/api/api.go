// Пакет api содержит структуры и функции, предназначенные для унификации ответов http-сервера, обработки запросов http-client.
package api

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Передаваемые статусы в ответе
const (
	StatusError string = "Error"
	StatusOK    string = "OK"
)

type Response struct {
	Status string
	Error  string
}

// Error формирует Response со статусом StatusError и описанием ошибки.
func Error(errDescription string) Response {
	return Response{
		Status: StatusError,
		Error:  errDescription,
	}
}

// Error формирует Response со статусом StatusOK.
func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

// ValidationError обрабатывает ошибки валидации, преобразует описание ошибок в читаемый вид.
func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}

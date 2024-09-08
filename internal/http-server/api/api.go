package api

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	StatusError string = "Error"
	StatusOK    string = "OK"
)

type Response struct {
	Status string
	Error  string
}

func Error(errDescription string) Response {
	return Response{
		Status: StatusError,
		Error:  errDescription,
	}
}

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

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

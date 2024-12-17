package types

import (
	"fmt"

	"github.com/DaviMoreira27/CalendarSync/internal/common/enums"
)

type Error interface {
	error
	Status() int
}

type HttpErrorType struct {
	Message    string
	StatusCode int
}

func (e HttpErrorType) Error() string {
	return e.Message
}

func (e HttpErrorType) Status() int {
	return e.StatusCode
}


type HttpOperation struct {
	Method enums.HttpMethod
	Operation string
}

type InternalError struct {
	Err error
}

func (internalError *InternalError) Error() string {
	return fmt.Sprintf("Status: 500 \nError: %v", internalError.Err)
}

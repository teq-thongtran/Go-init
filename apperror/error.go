package apperror

import (
	"errors"
	"fmt"
	"strings"
)

type TeqError struct {
	Raw       error
	ErrorCode string
	HTTPCode  int
	Message   string
	IsSentry  bool
}

func (e TeqError) Error() string {
	if e.Raw != nil {
		fmt.Println(e.Raw)
	}

	return e.Message
}

func (e TeqError) Is(target error) bool {
	if e.Raw != nil {
		return errors.Is(e.Raw, target)
	}

	return strings.Contains(e.Error(), target.Error())
}

func NewError(err error, httpCode int, errCode string, message string, isSentry bool) TeqError {
	return TeqError{
		Raw:       err,
		ErrorCode: errCode,
		HTTPCode:  httpCode,
		Message:   message,
		IsSentry:  isSentry,
	}
}

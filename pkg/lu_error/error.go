package lue

import (
	"fmt"
	"net/http"
)

type IError interface {
	SetHttpCode(code int) IError
	Error() string
	GetCode() string
	GetHTTP() int
}

type Error struct {
	error    any
	code     int
	httpCode int
}

func New(err any, code int) *Error {
	return &Error{
		error: err,
		code:  code,
	}
}

func (e *Error) SetHttpCode(code int) IError {
	e.httpCode = code
	return e
}

func (e *Error) Error() string {
	return fmt.Sprint(e.error)
}

func (e *Error) GetCode() string {
	return CodeInText[e.code]
}

func (e *Error) GetHTTP() int {
	if e.httpCode == 0 {
		e.httpCode = http.StatusInternalServerError
	}
	return e.httpCode
}

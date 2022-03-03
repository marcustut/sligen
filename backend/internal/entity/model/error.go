package model

import (
	"fmt"

	"github.com/marcustut/fyp/backend/internal/util/environment"

	"github.com/pkg/errors"
)

const (
	// DBError is error code of database
	DBError = "DB_ERROR"
	// GraphQLError is error code of graphql
	GraphQLError = "GRAPHQL_ERROR"
	// NotFoundError is error code of not found
	NotFoundError = "NOT_FOUND_ERROR"
	// ValidationError is error code of validation
	ValidationError = "VALIDATION_ERROR"
	// BadRequestError is error code of request
	BadRequestError = "BAD_REQUEST_ERROR"
	// InternalServerError is error code of server error
	InternalServerError = "INTERNAL_SERVER_ERROR"
)

// StackTrace is used to check to see if the error has already been wrapped by errors.WithStack
type StackTrace interface {
	StackTrace() errors.StackTrace
}

// Error is the standard error type
type Error interface {
	Error() string
	Code() string
	Extensions() map[string]interface{}
	Unwrap() error
}

// NewDBError returns error message related database
func NewDBError(e error) error {
	return newError(
		DBError,
		e.Error(),
		map[string]interface{}{
			"code": DBError,
		},
		e,
	)
}

// NewGraphQLError returns error message related graphql
func NewGraphQLError(e error) error {
	return newError(
		GraphQLError,
		e.Error(),
		map[string]interface{}{
			"code": GraphQLError,
		},
		e,
	)
}

// NewNotFoundError returns error message related not found
func NewNotFoundError(e error, value interface{}) error {
	return newError(
		NotFoundError,
		e.Error(),
		map[string]interface{}{
			"code":  NotFoundError,
			"value": value,
		},
		e,
	)
}

// NewInvalidParamError returns error message related param
func NewInvalidParamError(e error, value interface{}) error {
	return newError(
		BadRequestError,
		e.Error(),
		map[string]interface{}{
			"code":  BadRequestError,
			"value": value,
		},
		e,
	)
}

// NewInvalidRequestBodyError returns error message related request body
func NewInvalidRequestBodyError(e error) error {
	return newError(
		BadRequestError,
		e.Error(),
		map[string]interface{}{
			"code": BadRequestError,
		},
		e,
	)
}

// NewValidationError returns error message related validation
func NewValidationError(e error) error {
	return newError(
		ValidationError,
		e.Error(),
		map[string]interface{}{
			"code": ValidationError,
		},
		e,
	)
}

// NewInternalServerError returns error message related syntax or other issues
func NewInternalServerError(e error) error {
	return newError(
		InternalServerError,
		e.Error(),
		map[string]interface{}{
			"code": InternalServerError,
		},
		e,
	)
}

type err struct {
	err        error
	code       string
	message    string
	extensions map[string]interface{}
}

func (e *err) Error() string                      { return e.message }
func (e *err) Code() string                       { return e.code }
func (e *err) Extensions() map[string]interface{} { return e.extensions }
func (e *err) Unwrap() error                      { return e.err }

// IsStackTrace checks to see if the error with stack strace
func IsStackTrace(e error) bool {
	_, ok := e.(StackTrace)
	return ok
}

// IsError checks to see if the error is generated by model
func IsError(e error) bool {
	_, ok := e.(Error)
	return ok
}

func newError(code string, message string, extensions map[string]interface{}, e error) error {
	newErr := &err{
		err:        e,
		code:       code,
		message:    message,
		extensions: extensions,
	}
	if IsStackTrace(e) {
		return newErr
	}

	return withStackTrace(newErr)
}

func withStackTrace(e error) error {
	ews := errors.WithStack(e)

	if environment.IsDev() {
		fmt.Printf("%+v\n", ews)
	}

	return ews
}

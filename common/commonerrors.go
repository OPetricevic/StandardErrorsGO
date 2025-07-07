package common

import (
	"errors" // Import the errors package
	"fmt"
)

var (
	ErrNotFound             = errors.New("not found")
	ErrInternal             = errors.New("internal server error")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrForbidden            = errors.New("forbidden")
	ErrConflict             = errors.New("conflict")
	ErrTooManyRequests      = errors.New("too many requests")
	ErrBadRequest           = errors.New("bad request")
	ErrServiceUnavailable   = errors.New("service unavailable")
	ErrGatewayTimeout       = errors.New("gateway timeout")
	ErrMethodNotAllowed     = errors.New("method not allowed")
	ErrUnprocessableEntity  = errors.New("unprocessable entity")
	ErrUnsupportedMediaType = errors.New("unsupported media type")
	ErrPreconditionFailed   = errors.New("precondition failed")
	ErrExpectationFailed    = errors.New("expectation failed")
	ErrNotImplemented       = errors.New("not implemented")
	ErrBadGateway           = errors.New("bad gateway")
	ErrLengthRequired       = errors.New("length required")
	ErrInsufficientStorage  = errors.New("insufficient storage")
)

type ErrInvalidUUID struct {
	Value string
}

type ErrPathBodyMismatchedValue struct {
	PathFieldName string
	PathValue     any
	BodyValue     any
}

func (e ErrInvalidUUID) Error() string {
	return fmt.Sprintf("the value %q is an invalid UUID", e.Value)
}

func (e ErrPathBodyMismatchedValue) Error() string {
	return fmt.Sprintf("the value for the path parameter %q does not match the value from the body (%v != %v)", e.PathFieldName, e.PathValue, e.BodyValue)
}

// ErrUUIDValidationFailed is a reusable error for when UUID validation fails for multiple fields.
type ErrUUIDValidationFailed struct {
	InvalidUUIDs map[string]string
}

func (e ErrUUIDValidationFailed) Error() string {
	var result string
	for field, uuid := range e.InvalidUUIDs {
		result += fmt.Sprintf("invalid UUID for %s: %q; ", field, uuid)
	}
	return result
}

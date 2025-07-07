package repositoryerrors

import (
	"errors"
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

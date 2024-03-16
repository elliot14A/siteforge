package errors

import "errors"

type CustomError error

var (
	ErrNotFound        CustomError = errors.New("not found")
	ErrInvalidPassword CustomError = errors.New("invalid password")
	ErrInternalError   CustomError = errors.New("internal error")
)

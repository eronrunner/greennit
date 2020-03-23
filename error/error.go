package error

import "errors"

// AppError - an abstraction represents for all App errors
type AppError struct {
	Error   error
	Message string
	Code    int
}

var MsgBadRequest = "Bad request"

var MsgInvalidInput = "Invalid input"

var ErrObjectExist = errors.New("Object has already existed")

var ErrObjectNotExist = errors.New("Object hasn't already existed")

var ErrBadCredential = errors.New("Bad credential")

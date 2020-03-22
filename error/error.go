package error

import "errors"

// AppError - an abstraction represents for all App errors
type AppError struct {
	Error   error
  Message string
	Code    int
}

var ErrObjectExist = errors.New("Object has already existed")
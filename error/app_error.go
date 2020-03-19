package error

// AppError - an abstraction represents for all App errors
type AppError struct {
	Error   error
  Message string
	Code    int
}
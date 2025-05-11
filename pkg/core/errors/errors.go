package errors

type ShynobiError interface {
	ErrorMessage() string
	ErrorCode() int
}

type SimpleShynobiError struct {
	Message string
	Code    int
}

func FromError(err error) *SimpleShynobiError {
	return &SimpleShynobiError{
		Message: err.Error(),
		Code:    500,
	}
}

func (e SimpleShynobiError) ErrorMessage() string {
	return e.Message
}

func (e SimpleShynobiError) ErrorCode() int {
	return e.Code
}

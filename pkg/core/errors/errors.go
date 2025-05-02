package errors

type ShynobiError interface {
	ErrorMessage() string
	ErrorCode() int
}

type SimpleShynobiError struct {
	Message string
	Code    int
}

func (e SimpleShynobiError) ErrorMessage() string {
	return e.Message
}

func (e SimpleShynobiError) ErrorCode() int {
	return e.Code
}

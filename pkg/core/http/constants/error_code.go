package constants

type ErrorCodeConstant struct{}

var (
	ErrorCodeConstants = ErrorCodeConstant{}
)

func (e ErrorCodeConstant) Success() int       { return 0 }
func (e ErrorCodeConstant) Unimplemented() int { return -1 }

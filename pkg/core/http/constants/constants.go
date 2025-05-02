package constants

type MessageIdConstant struct{}
type ErrorCodeConstant struct{}

var (
	MessageIdConstants = MessageIdConstant{}
	ErrorCodeConstants = ErrorCodeConstant{}
)

// Error Code
func (e ErrorCodeConstant) Success() int       { return 0 }
func (e ErrorCodeConstant) Unimplemented() int { return -1 }

// Message Id
func (m MessageIdConstant) Unknown() int { return 0 }

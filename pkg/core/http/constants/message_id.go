package constants

type MessageIdConstant struct{}

var (
	MessageIdConstants = MessageIdConstant{}
)

func (m MessageIdConstant) Unknown() int { return 0 }

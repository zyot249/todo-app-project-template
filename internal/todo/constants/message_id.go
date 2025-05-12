package constants

import (
	"todo-app/pkg/core/http/constants"
)

type TodoMessageIdConstant struct {
	constants.MessageIdConstant
}

var TodoMessageIdConstants = TodoMessageIdConstant{}

func (m TodoMessageIdConstant) GetTodo() int    { return 1000 }
func (m TodoMessageIdConstant) CreateTodo() int { return 1001 }
func (m TodoMessageIdConstant) UpdateTodo() int { return 1002 }
func (m TodoMessageIdConstant) DeleteTodo() int { return 1003 }

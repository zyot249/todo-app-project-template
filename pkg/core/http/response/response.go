package response

import (
	"net/http"
	"todo-app/pkg/core/errors"

	"github.com/go-chi/render"
)

type Response struct {
	MessageId int `json:"messageId"`
	ErrorCode int `json:"errorCode"`
	Data      any `json:"data"`
}

func RespondSuccess(w http.ResponseWriter, r *http.Request, messageId int, data interface{}) {
	render.Status(r, http.StatusOK)
	render.Respond(w, r, Response{
		MessageId: messageId,
		ErrorCode: 0,
		Data:      data,
	})
}

func RespondError(w http.ResponseWriter, r *http.Request, messageId int, err errors.ShynobiError) {
	render.Status(r, http.StatusBadRequest)
	render.Respond(w, r, Response{
		MessageId: messageId,
		ErrorCode: err.ErrorCode(),
		Data:      err.ErrorMessage(),
	})
}

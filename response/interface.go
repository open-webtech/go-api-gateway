package response

import (
	"net/http"

	"github.com/secondtruth/go-api-gateway/response/format"
)

type HttpResponder interface {
	SetFormatter(formatter format.HttpResponseFormatter)
	Send(w http.ResponseWriter, content []byte, code int) error
	SendMsg(w http.ResponseWriter, text, caption string, code int) error
	SendData(w http.ResponseWriter, data any) error
	SendClientError(w http.ResponseWriter, msg string, code int) error
	SendServerError(w http.ResponseWriter, msg string, code int) error
}

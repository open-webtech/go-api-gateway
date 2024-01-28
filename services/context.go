package services

import (
	"github.com/secondtruth/go-api-gateway/logging"
	"github.com/secondtruth/go-api-gateway/response"
)

type Context struct {
	Responder response.HttpResponder
	AccessLog logging.AccessLogger
}

func NewContextWithDefaults() Context {
	return Context{
		Responder: response.NewGenericResponder(nil),
	}
}

package gateway

import (
	"net/http"

	"github.com/secondtruth/go-api-gateway/handlers"
	"github.com/secondtruth/go-api-gateway/services"
	reverseproxy "github.com/secondtruth/go-reverse-proxy"
)

type HandlerFactory struct {
	Transport      http.RoundTripper
	RequestHeader  http.Header
	ModifyResponse func(*http.Response) error
	Services       services.Context
}

func NewHandlerFactory(sc services.Context) *HandlerFactory {
	return &HandlerFactory{
		Services: sc,
	}
}

func (f *HandlerFactory) MakeReverseProxy(remote string) (*reverseproxy.ReverseProxyMux, error) {
	rp, err := handlers.ReverseProxy(remote, f.Services)
	rp.Transport = f.Transport
	rp.RequestHeader = f.RequestHeader
	rp.ModifyResponse = f.ModifyResponse
	return rp, err
}

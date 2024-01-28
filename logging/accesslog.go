package logging

import (
	"io"
	"log"
	"net/http"
)

type SimpleAccessLog struct {
	handler log.Logger
}

func NewSimpleAccessLog(out io.Writer, prefix string, flag int) *SimpleAccessLog {
	return &SimpleAccessLog{
		handler: *log.New(out, prefix, flag),
	}
}

func (l *SimpleAccessLog) LogAccess(r *http.Request) {
	l.handler.Printf("[%s] %s %s", r.Host, r.Method, r.URL.Path)
}

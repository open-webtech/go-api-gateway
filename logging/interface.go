package logging

import "net/http"

type AccessLogger interface {
	LogAccess(request *http.Request)
}

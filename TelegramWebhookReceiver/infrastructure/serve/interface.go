package serve

import "net/http"

type Server interface {
	Start() error
	Stop() error
	Restart() error

	SetRouter() error
	Configure(options map[string]interface{}) error
}

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

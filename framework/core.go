package framework

import "net/http"

// Framework core struct
type Core struct {
}

// Init framework core
func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}

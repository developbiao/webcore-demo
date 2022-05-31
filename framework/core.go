package framework

import (
	"log"
	"net/http"
	"strings"
)

// Framework core struct
type Core struct {
	// Level 2 map
	router map[string]map[string]ControllerHandler
}

// Init framework core
func NewCore() *Core {
	// Define level 2 map
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}

	// Put level 2 routers to level 1 map
	router := map[string]map[string]ControllerHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter

	return &Core{router: router}

}

// Get method
func (c *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
}

// Post method
func (c *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}

// Put method
func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}

// Delete method
func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
}

// Find route, if not found, return nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// uri and method must be convert to uppercase
	uri := request.URL.Path
	method := request.Method
	upperUri := strings.ToUpper(uri)
	upperMethod := strings.ToUpper(method)

	// Find from level 1 map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		// Find from level 2 map
		if handler, ok := methodHandlers[upperUri]; ok {
			return handler
		}
	}
	return nil
}

// Serve
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(request, response)

	// Find route
	router := c.FindRouteByRequest(request)
	if router == nil {
		ctx.Json(404, "Not found")
		return
	}

	// Call route
	if err := router(ctx); err != nil {
		ctx.Json(500, "Internal server error")

	}
	log.Println("core.router")

	router(ctx)
}

// Group http wrapper
func (c *Core) Group(prefix string) *Group {
	return NewGroup(c, prefix)
}

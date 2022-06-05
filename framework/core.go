package framework

import (
	"log"
	"net/http"
	"strings"
)

// Framework core struct
type Core struct {
	router map[string]*Tree
}

// Init framework core
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{router: router}

}

// Get method
func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Post method
func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Put method
func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Delete method
func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Find route, if not found, return nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// uri and method must be convert to uppercase
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// Find from level 1 map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
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

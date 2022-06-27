package framework

import (
	"log"
	"net/http"
	"strings"
)

// Framework core struct
type Core struct {
	router      map[string]*Tree    // all routers
	middlewares []ControllerHandler // set middlewares
}

// Init framework core
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	core := &Core{router: router}
	return core
}

// Register middleware
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

// Get method
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Post method
func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Put method
func (c *Core) Put(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Delete method
func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error:", err)
	}
}

// Find route, if not found, return nil
func (c *Core) FindRouteByRequest(request *http.Request) *node {
	// uri and method must be convert to uppercase
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// Find from level 1 map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.root.matchNode(uri)
	}
	return nil
}

// Serve
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(request, response)

	// Find route
	node := c.FindRouteByRequest(request)
	if node == nil {
		// Not found route
		ctx.SetStatus(404).Json("Not found")
		return
	}

	// Set context handlers
	ctx.SetHandlers(node.handlers)

	// set route paramters
	params := node.parseParamsFromEndNode(request.URL.Path)
	ctx.SetParams(params)

	// Call route
	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("Internal server error")
		return
	}
}

// Group http wrapper
func (c *Core) Group(prefix string) *Group {
	return NewGroup(c, prefix)
}

package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context

	// Flag is timeout
	hasTimeout bool
	// Write safe
	writerMux *sync.Mutex

	// Current request handler chain
	handlers []ControllerHandler
	// Current coursor index
	index int

	params map[string]string // url route params
}

// New Context
func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
		index:          -1,
	}
}

// Writer Mux
func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

// Get request
func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

// Get response
func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

// Set timeout
func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

// Check is timeout
func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// Base context from request
func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// Done
func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

// Implement deadline
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

// Implement error
func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

// #End Form methods

// Set handlers on context
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

// Set params
func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

// Next
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

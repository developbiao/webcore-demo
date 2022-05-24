package framework

import (
	"context"
	"net/http"
	"strconv"
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
}

// New Context
func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
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

// #Begin Query methods

// Region query url
func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			intval, err := strconv.Atoi(vals[len-1])
			if err != nil {
				return intval
			}
			return intval
		}
	}
	return def
}

// Query string
func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
	}
	return def
}

// Query array
func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

// Query all
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}
}

// #End Query methods

// #Begin Form methods

// FormInt
func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			intval, err := strconv.Atoi(vals[len-1])
			if err != nil {
				return intval
			}
			return intval
		}
	}
	return def
}

// FormString
func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

// FormAll
func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.Form)
	}
	return map[string][]string{}
}

package gin

import (
	"context"

	"github.com/developbiao/webcore-demo/framework"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

// Bind engin implemention bind
func (engin *Engine) Bind(provider framework.ServiceProvider) error {
	return engin.container.Bind(provider)
}

// IsBind engin implemnetation  is bind
func (engin *Engine) IsBind(key string) bool {
	return engin.container.IsBind(key)
}

// ----------------------
// Context impelmentation

// Make context implementation make
func (ctx *Context) Make(key string) (interface{}, error) {
	return ctx.container.Make(key)
}

// MustMake context implemnetation must make
func (ctx *Context) MustMake(key string) interface{} {
	return ctx.container.MustMake(key)
}

// MakeNew context implemnetation make new
func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return ctx.container.MakeNew(key, params)
}

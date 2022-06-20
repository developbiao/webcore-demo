package framework

import (
	"github.com/spf13/cast"
)

type IRequest interface {
	// Request url parameters
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) interface{}

	// Match in route parmeters
	// e.g /user/:id
	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamBool(key string, def bool) (bool, bool)

	// Get parameters from form data
	FromInt(key string, def int) (int, bool)
	FromInt64(key string, def int64) (int64, bool)
	FromFloat64(key string, def float64) (float64, bool)
	FromFloat32(key string, def float32) (float32, bool)
	FromString(key string, def string) (string, bool)
	FromStringSlice(key string, def []string) ([]string, bool)
	FromFile(key string) interface{}

	// json body
	BindJson(obj interface{}) error

	// xml body
	BindXml(obj interface{}) error

	// Other body
	GetRawData(obj interface{}) error

	// Base information
	Uri() string
	Method() string
	Host() string
	ClientIP() string

	// headers
	Headers() map[string][]string
	Header(key string) (string, bool)

	// Cookies
	Cookies() map[string][]string
	Cookie(key string) (string, bool)
}

// Get request uri all parameters
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

// #Begin Query methods

// Region query url
func (ctx *Context) QueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

// Query int64
func (ctx *Context) QueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

// Query float64
func (ctx *Context) QueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.toFloat64(vals[0]), true
		}
	}
	return def, false
}

// Query float32
func (ctx *Context) QueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

// Query bool
func (ctx *Context) QueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

// QueryString
func (ctx *Context) QueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.Tostring(vals[0]), true
		}
	}
	return def, false
}

// Query string slice
func (ctx *Context) QueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals, true
		}
	}
	return def, false
}

// Query
func (ctx *Context) Query(key string) interface{} {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return nil
}

// #End Query methods

// TODO:: Implement Params

// #Begin Form methods

// FormInt
func (ctx *Context) FormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

// FormString
func (ctx *Context) FormString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals[0], true
	}
	return def, false
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

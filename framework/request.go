package framework

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"mime/multipart"

	"github.com/spf13/cast"
)

const defaultMultipartMemory = 32 << 20 // 32MB

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
	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormFloat32(key string, def float32) (float32, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) ([]string, bool)
	FormFile(key string) (*multipart.FileHeader, error)
	Form(key string) interface{}

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
	ClientIp() string

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
			return cast.ToFloat64(vals[0]), true
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
			return cast.ToString(vals[0]), true
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

// FormInt64
func (ctx *Context) FormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

// FormFloat64
func (ctx *Context) FormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

// FormFloat32
func (ctx *Context) FormFloat32(key string, def float32) (float32, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

// FormBool
func (ctx *Context) FormFloatBool(key string, def bool) (bool, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return cast.ToBool(vals[0]), true
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

// FromStringSlice
func (ctx *Context) FromStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals, true
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

// FromFile
func (ctx *Context) FormFile(key string) (*multipart.FileHeader, error) {
	if ctx.request.MultipartForm == nil {
		if err := ctx.request.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := ctx.request.FormFile(key)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}

// From
func (ctx *Context) Form(key string) interface{} {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return nil
}

// BindJson
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// BindXml
func (ctx *Context) BindXml(obj interface{}) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		// 重新填充request.Body，为后续的逻辑二次读取做准备
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// Parsing to obj structure
		err = xml.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// GetRawData
func (ctx *Context) GetRawData() ([]byte, error) {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return nil, err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return body, nil
	}
	return nil, errors.New("ctx.request empty")
}

// Uri
func (ctx *Context) Uri() string {
	return ctx.request.RequestURI
}

// Method
func (ctx *Context) Method() string {
	return ctx.request.Method
}

// ClientIp
func (ctx *Context) ClientIp() string {
	r := ctx.request
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}

// Headers
func (ctx *Context) Headers() map[string][]string {
	return ctx.request.Header
}

// Header
func (ctx *Context) Header(key string) (string, bool) {
	vals := ctx.request.Header.Values(key)
	if vals == nil || len(vals) <= 0 {
		return "", false
	}
	return vals[0], true
}

// Cookies
func (ctx *Context) Cookies() map[string]string {
	cookies := ctx.request.Cookies()
	ret := map[string]string{}
	for _, cookie := range cookies {
		ret[cookie.Name] = cookie.Value
	}
	return ret
}

// Cookie
func (ctx *Context) Cookie(key string) (string, bool) {
	cookies := ctx.Cookies()
	if val, ok := cookies[key]; ok {
		return val, true
	}
	return "", false
}

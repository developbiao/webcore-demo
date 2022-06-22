package framework

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type IResponse interface {
	// Json output
	Json(obj interface{}) IResponse

	// Jsonp output
	Jsonp(obj interface{}) IResponse

	// xml output
	Xml(file string, obj interface{}) IResponse

	// html output
	Html(obj interface{}) IResponse

	// Text
	Text(format string, values ...interface{}) IResponse

	// Redirect
	Redirect(path string) IResponse

	// Header
	SetHeader(key string, val string) IResponse

	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// Set status code
	SetStatus(code int) IResponse

	// Set 200 ok status
	SetOkStatus() IResponse
}

// Set header
func (ctx *Context) SetHeader(key string, val string) *Context {
	ctx.responseWriter.Header().Add(key, val)
	return ctx
}

// SetCookie
func (ctx *Context) SetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) *Context {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return ctx
}

// Set status code
func (ctx *Context) SetStatus(code int) *Context {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

func (ctx *Context) SetOkStatus() *Context {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}

// Json
func (ctx *Context) Json(obj interface{}) *Context {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.responseWriter.Write(byt)
	return ctx
}

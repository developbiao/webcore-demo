package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
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

// Jsonp output
func (ctx *Context) Jsonp(obj interface{}) *Context {
	// Get request callback parameters
	callbackFunc, _ := ctx.QueryString("callback", "callback_function")
	ctx.SetHeader("Context-Type", "application/javascript")
	// Prevent attack XSS
	callback := template.JSEscapeString(callbackFunc)

	// Output function
	_, err := ctx.responseWriter.Write([]byte(callback))
	if err != nil {
		return ctx
	}

	// Output left parenthesis
	_, err = ctx.responseWriter.Write([]byte("("))
	if err != nil {
		return ctx
	}

	// Fcuntion paramters
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		return ctx
	}

	// Output right parenthesis
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}
	return ctx

}

// Xml output
func (ctx *Context) Xml(obj interface{}) *Context {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
		ctx.SetHeader("Content-Type", "application/html")
	}
	ctx.SetHeader("Content-Type", "application/html")
	ctx.responseWriter.Write(byt)
	return ctx
}

// Html output
func (ctx *Context) Html(file string, obj interface{}) *Context {
	// Read template file create template instance
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}

	// Execute obj
	if err := t.Execute(ctx.responseWriter, obj); err != nil {
		return ctx
	}

	ctx.SetHeader("Context-Type", "application/html")
	return ctx
}

// Text
func (ctx *Context) Text(format string, values ...interface{}) *Context {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Content-Type", "application/text")
	ctx.responseWriter.Write([]byte(out))
	return ctx
}

// Redirect
func (ctx *Context) Redirect(path string) *Context {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}

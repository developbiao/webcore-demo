package http

import "github.com/developbiao/webcore-demo/framework/gin"

// NewHttpEngine is command
func NewHttpEngine() (*gin.Engine, error) {
	// Set release, for defrault prevent ouput debug information
	gin.SetMode(gin.ReleaseMode)
	// Startup default web engine
	r := gin.Default()

	// Binding business route
	Routes(r)
	return r, nil
}

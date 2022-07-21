package main

import (
	"net/http"

	"github.com/developbiao/webcore-demo/framework"
	"github.com/developbiao/webcore-demo/framework/provider/app"
	"github.com/developbiao/webcore-demo/framework/provider/kernel"
)

func main() {
	// Initlization container
	container := framework.NewWebContainer()

	// Bind app service provider
	container.Bind(&app.WebAppProvider{})

	// Initilization http engine
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.WebKernelProvider{HttpEngine: engine})
	}

	// Run room command
	console.RunCommand(container)
}

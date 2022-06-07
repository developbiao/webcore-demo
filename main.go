package main

import (
	"log"
	"net/http"

	"github.com/developbiao/webcore-demo/framework"
	"github.com/developbiao/webcore-demo/middleware"
)

func main() {
	core := framework.NewCore()
	core.Use(
		middleware.Test1(),
		middleware.Test2(),
		middleware.Test3(),
	)
	registerRouter(core)
	server := &http.Server{
		// Customer request core handler
		Handler: core,
		Addr:    ":8080",
	}
	log.Println("Server on :8080")
	server.ListenAndServe()
}

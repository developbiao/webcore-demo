package main

import (
	"log"
	"net/http"

	"github.com/developbiao/webcore-demo/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		// Customer request core handler
		Handler: core,
		Addr:    ":8080",
	}
	log.Println("Server on :8080")
	server.ListenAndServe()
}

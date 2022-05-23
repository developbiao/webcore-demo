package main

import (
	"net/http"

	"github.com/developbiao/webcore-demo/framework"
)

func main() {
	//
	server := &http.Server{
		// Customer request core handler
		Handler: framework.NewCore(),
		Addr:    ":8080",
	}
	server.ListenAndServe()
}

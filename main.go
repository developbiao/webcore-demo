package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/developbiao/webcore-demo/framework"
	"github.com/developbiao/webcore-demo/middleware"
)

func main() {
	core := framework.NewCore()
	// core.Use(
	// 	middleware.Test1(),
	// 	middleware.Test2(),
	// )

	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)
	server := &http.Server{
		// Customer request core handler
		Handler: core,
		Addr:    ":8080",
	}

	// Serve with goroutine
	go func() {
		log.Println("Server on :8080")
		server.ListenAndServe()
	}()

	// Current goroutine signal
	quit := make(chan os.Signal)

	// Monitor signal: SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// Block goroutine wait signal
	<-quit

	// Call Server.Shutdown graceful
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

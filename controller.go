package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/developbiao/webcore-demo/framework"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	// Wild go routine
	go func() {
		time.Sleep(10 * time.Second)
		c.SetStatus(200).Json("ok")

		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.SetStatus(500).Json("panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.SetStatus(500).Json("time out")
		c.SetHasTimeout()
	}

	return nil
}

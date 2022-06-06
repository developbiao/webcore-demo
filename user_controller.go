package main

import (
	"time"

	"github.com/developbiao/webcore-demo/framework"
)

func UserLoginController(c *framework.Context) error {
	time.Sleep(time.Second * 2)
	c.Json(200, "ok, UserLoginController")
	return nil
}

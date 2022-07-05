package main

import (
	"time"

	"github.com/developbiao/webcore-demo/framework"
)

func UserLoginController(c *framework.Context) {
	foo, _ := c.QueryString("foo", "def")
	// Wait 10 seconds end
	time.Sleep(time.Second * 10)
	c.SetStatus(200).Json("ok, UserLoginController: " + foo)
}

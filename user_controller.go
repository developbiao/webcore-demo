package main

import (
	"time"

	"github.com/developbiao/webcore-demo/framework/gin"
)

func UserLoginController(c *gin.Context) {
	foo, _ := c.DefaultQueryString("foo", "def")
	// Wait 10 seconds end
	time.Sleep(time.Second * 10)
	c.ISetOkStatus().IJson("ok, UserLoginController: " + foo)
}

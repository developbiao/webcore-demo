package main

import "github.com/developbiao/webcore-demo/framework"

func registerRouter(core *framework.Core) {
	core.Get("/foo", FooControllerHandler)
	core.Get("/user/login", UserLoginControllerHandler)
}

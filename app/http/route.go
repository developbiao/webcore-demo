package http

import (
	"github.com/developbiao/webcore-demo/app/http/module/demo"
	"github.com/developbiao/webcore-demo/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")

	demo.Register(r)
}

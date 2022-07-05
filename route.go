package main

import (
	"github.com/developbiao/webcore-demo/framework/gin"
	"github.com/developbiao/webcore-demo/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	core.GET("/user/login", middleware.Test3(), UserLoginController)

	// Group prefix routes
	subjectApi := core.Group("/subject")
	{
		// Apply middleware test1
		subjectApi.DELETE("/:id", middleware.Test1(), SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)

		// Embedding group
		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}
}

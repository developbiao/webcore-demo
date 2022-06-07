package main

import (
	"github.com/developbiao/webcore-demo/framework"
	"github.com/developbiao/webcore-demo/middleware"
)

func registerRouter(core *framework.Core) {
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// Group prefix routes
	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		// Embedding group
		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}

}

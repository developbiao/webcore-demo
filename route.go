package main

import "github.com/developbiao/webcore-demo/framework"

func registerRouter(core *framework.Core) {
	core.Get("/user/login", UserLoginController)

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

package main

import (
	"fmt"

	"github.com/developbiao/webcore-demo/framework/gin"
	"github.com/developbiao/webcore-demo/provider/demo"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectAddController")
}

// corresponding route /subject/list/all
func SubjectListController(c *gin.Context) {
	// Get demo instance
	demoService := c.MustMake(demo.Key).(demo.Service)

	//  Callback method on instance
	foo := demoService.GetFoo()

	// Output result
	c.ISetOkStatus().IJson(foo)
	// c.ISetOkStatus().IJson("ok, SubjectListController")
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("ok, SubjectGetController:" + fmt.Sprint(subjectId))

}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok, SubjectNameController")
}

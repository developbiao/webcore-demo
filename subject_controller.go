package main

import "github.com/developbiao/webcore-demo/framework"

// SubjectAddController
func SubjectAddController(c *framework.Context) error {
	c.Json(200, "ok, SubjectAddController")
	return nil
}

// SubjectListController
func SubjectListController(c *framework.Context) error {
	c.Json(200, "ok, SubjectListController")
	return nil
}

// SubjectDelController
func SubjectDelController(c *framework.Context) error {
	c.Json(200, "ok, SubjectDelController")
	return nil
}

// SubjectUpdateController
func SubjectUpdateController(c *framework.Context) error {
	c.Json(200, "ok, SubjectUpdateController")
	return nil
}

// SubjectGetController
func SubjectGetController(c *framework.Context) error {
	c.Json(200, "ok, SubjectGetController")
	return nil
}

// SubjectNameController
func SubjectNameController(c *framework.Context) error {
	c.Json(200, "ok, SubjectNameController")
	return nil
}

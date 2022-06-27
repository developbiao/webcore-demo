package main

import (
	"time"

	"github.com/developbiao/webcore-demo/framework"
)

// SubjectAddController
func SubjectAddController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectAddController")
	return nil
}

// SubjectListController
func SubjectListController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectListController")
	return nil
}

// SubjectDelController
func SubjectDelController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectDelController")
	return nil
}

// SubjectUpdateController
func SubjectUpdateController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectUpdateController")
	return nil
}

// SubjectGetController
func SubjectGetController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectGetController")
	return nil
}

// SubjectNameController
func SubjectNameController(c *framework.Context) error {
	time.Sleep(time.Second * 2)
	c.SetStatus(200).Json("ok, SubjectNameController")
	return nil
}

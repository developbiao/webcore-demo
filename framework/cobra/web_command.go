package cobra

import "github.com/developbiao/webcore-demo/framework"

func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

func (c *Command) GetContainer() framework.Container {
	return c.container
}
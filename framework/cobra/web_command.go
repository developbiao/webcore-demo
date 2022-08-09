package cobra

import "github.com/developbiao/webcore-demo/framework"

// Set container
func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

// Get container
func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}

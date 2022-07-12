package demo

import "github.com/developbiao/webcore-demo/framework"

// DemoProvider
type DemoProvider struct {
	framework.ServiceProvider

	c framework.Container
}

// Implementation name identify
func (sp *DemoProvider) Name() string {
	return DemoKey
}

// Implementation Register
func (sp *DemoProvider) Register(c framework.Container) framework.NewInstance {
	return NewService
}

// Implementation IsDefer
func (sp *DemoProvider) IsDefer() bool {
	return false
}

// Implementation params
func (sp *DemoProvider) Params(c framework.Container) []interface{} {
	return []interface{}{sp.c}
}

// Implementation boot
func (sp *DemoProvider) Boot(c framework.Container) error {
	sp.c = c
	return nil
}

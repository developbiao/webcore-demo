package demo

import (
	"fmt"

	"github.com/developbiao/webcore-demo/framework"
)

// 存放服务提供方 ServiceProvider 的实现
// Service Provider
type DemoServiceProvider struct {
}

// Name 方法直接将服务对应的字符串凭证返回，这里返回 “web.demo”
func (sp *DemoServiceProvider) Name() string {
	return Key
}

// Register 方法是注册初始化服务实例的方法, 我们这里先暂定为NewDemoSevice
func (sp *DemoServiceProvider) Register(c framework.Container) framework.NewInstance {
	return newDemoService
}

// IsDefer 方法表示是否延迟实例化，我们这里设置为true,
// 将这个服务的实例化延迟到第一次make的时候
func (sp *DemoServiceProvider) IsDefer() bool {
	return true
}

// Params 方法表示实例的参数，我们这里只实例化一个参数：container
// 表示我们在 NewDemoServie 这个函数中， 只有一个参数，container
func (sp *DemoServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

// Boot 方法我们这里我们什么逻辑都不执行，只打印一行日志
func (sp *DemoServiceProvider) Boot(c framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}

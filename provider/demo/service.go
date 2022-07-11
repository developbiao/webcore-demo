package demo

import (
	"fmt"

	"github.com/developbiao/webcore-demo/framework"
)

// 服务具体的实现
type DemoService struct {
	// 实现接口
	Service

	// 参数
	c framework.Container
}

func newDemoService(params ...interface{}) (interface{}, error) {
	// 这里需要将参数展开
	c := params[0].(framework.Container)

	fmt.Println("new demo service")

	// 返回实例
	return &DemoService{c: c}, nil
}

// 实现接口
func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "I am first web service foo",
	}
}

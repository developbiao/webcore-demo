package demo

// 用于存放服务的接口文件和服务凭证

// Demo server key
const Key = "web:demo"

// Demo servcie interface
type Service interface {
	GetFoo() Foo
}

// Demo interface define data structure
type Foo struct {
	Name string
}

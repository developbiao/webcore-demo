package demo

const DemoKey = "demo"

// Service interface
type IService interface {
	GetAllStudent() []Student
}

// Student struct
type Student struct {
	ID   int
	Name string
}

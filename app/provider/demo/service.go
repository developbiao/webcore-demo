package demo

import "github.com/developbiao/webcore-demo/framework"

// Define Service
type Service struct {
	container framework.Container
}

// NewService
func NewService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	return &Service{container: container}, nil
}

// Implementation GetAllStudent
func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "Lisa",
		},
		{
			ID:   2,
			Name: "Jsona",
		},
		{
			ID:   3,
			Name: "Lucky",
		},
	}
}

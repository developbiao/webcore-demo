package kernel

import (
	"net/http"

	"github.com/developbiao/webcore-demo/framework/gin"
)

type WebKernelService struct {
	engine *gin.Engine
}

// NewWebkernelService
func NewWebKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &WebKernelService{engine: httpEngine}, nil
}

// HttpEngine
func (s *WebKernelService) HttpEngine() http.Handler {
	return s.engine
}

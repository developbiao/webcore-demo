package kernel

import (
	"github.com/developbiao/webcore-demo/framework"
	"github.com/developbiao/webcore-demo/framework/gin"
)

// WebKernelProvider web engine
type WebKernelProvider struct {
	HttpEngine *gin.Engine
}

// Register
func (provider *WebKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewWebKernelService
}

// Boot
func (provider *WebKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

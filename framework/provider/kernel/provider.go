package kernel

import (
	"github.com/developbiao/webcore-demo/framework"
	"github.com/developbiao/webcore-demo/framework/contract"
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

// Boot boot startup check is registration Engine, if not register new engine instance
func (provider *WebKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

// IsDefer  we want to initialize right from the start
func (provider *WebKernelProvider) IsDefer() bool {
	return false
}

// Params just a httpEngine
func (provider *WebKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}

}

// Name provider identify
func (provider *WebKernelProvider) Name() string {
	return contract.KernelKey
}

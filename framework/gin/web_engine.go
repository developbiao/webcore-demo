package gin

import "github.com/developbiao/webcore-demo/framework"

func (engine *Engine) SetContainer(container framework.Container) {
	engine.container = container
}

// Bind
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

// IsBind
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

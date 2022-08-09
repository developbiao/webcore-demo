package app

import (
	"github.com/developbiao/webcore-demo/framework"
	"github.com/developbiao/webcore-demo/framework/contract"
)

// WebAppProvider
type WebAppProvider struct {
	BaseFolder string
}

// Register register web app provider
func (w *WebAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewWebApp
}

// Boot boot invock
func (w *WebAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer is delay call
func (w *WebAppProvider) IsDefer() bool {
	return false
}

// Params get parameters
func (w *WebAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, w.BaseFolder}
}

// Name get identify
func (w *WebAppProvider) Name() string {
	return contract.AppKey
}

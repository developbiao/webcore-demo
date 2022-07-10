package framework

import (
	"errors"
	"fmt"
	"sync"
)

// Define container is server container
// Privoder bindig service
type Container interface {
	// Bind a service provider if keyword exists will be replace and return error
	Bind(provider ServiceProvider) error

	// IsBind identify check service provider has been bind
	IsBind(key string) bool

	// Make service by key identify
	Make(key string) (interface{}, error)

	// Make service by key identify, if key does not bound service provider will throw panic
	// So using the interfae must be sure keep service container already assign key identify bind service provider
	MustMake(key string) interface{}

	// MakeNew get service by key identify, but is not singleton
	// New instance by service provider register boot function pass paramters
	// This function is useful when you need to start different instances for different parameters
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// WebContainer
type WebContainer struct {
	Container

	// Providers store service provider, key is identify
	providers map[string]ServiceProvider

	// Store specific instances
	instances map[string]interface{}

	// Lock container operator
	lock sync.RWMutex
}

// NewWebContainer
func NewWebContainer() *WebContainer {
	return &WebContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintPrivoders print registered provider key name
func (web *WebContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range web.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

func (web *WebContainer) Bind(provider ServiceProvider) error {
	web.lock.Lock()
	defer web.lock.Unlock()

	key := provider.Name()
	web.providers[key] = provider

	// If provider is not defer
	if !provider.IsDefer() {
		if err := provider.Boot(web); err != nil {
			return err
		}

		// Instance
		params := provider.Params(web)
		method := provider.Register(web)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		web.instances[key] = instance
	}
	return nil
}

// IsBind implement
func (web *WebContainer) IsBind(key string) bool {
	return web.findServiceProvider(key) != nil
}

// findServiceProvicer
func (web *WebContainer) findServiceProvider(key string) ServiceProvider {
	web.lock.RLock()
	defer web.lock.Unlock()
	if sp, ok := web.providers[key]; ok {
		return sp
	}
	return nil
}

// Make
func (web *WebContainer) Make(key string) (interface{}, error) {
	return web.make(key, nil, false)
}

// MustMake
func (web *WebContainer) MustMake(key string) interface{} {
	serv, err := web.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

// MakeNew
func (web *WebContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return web.make(key, params, true)
}

// newInstance
func (web *WebContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(web); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(web)
	}
	method := sp.Register(web)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

// make real speific instance
func (web *WebContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	web.lock.RLock()
	defer web.lock.Unlock()
	// Query is already registered servcie proviceer, if not regisger will be return error
	sp := web.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return web.newInstance(sp, params)
	}

	// If don't need force new instance on container already instanced
	// will straight retrive instance
	if ins, ok := web.instances[key]; ok {
		return ins, nil
	}

	// If container not instance will be once instance
	inst, err := web.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}
	web.instances[key] = inst
	return inst, nil
}

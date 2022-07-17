package app

import (
	"errors"
	"path/filepath"
	"github.com/developbiao/webcore-demo/framework"
	"github.com/developbiao/webcore-demo/framework/util"
	"flag"
)

// WebApp reprenset web framework App implement
type WebApp struct {
	container framework.Container 	// Service container
	baseFolder string 				// Base folder path
}


// Version implement
func (w WebApp) Version() string {
	return "0.0.1"
}

// BaseFolder implement
func (w WebApp) BaseFolder() string {
	if w.baseFolder != "" {
		return w.baseFolder
	}

	// If not set path otherwise using parmater
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder paramter, default is current path")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	// If not sepcifiy paramter use current path
	return util.GetExecDirectory()
}

// ConfigFolder implement
func (w WebApp) ConfigFolder() string {
	return filepath.Join(w.BaseFolder(), "config")
}

// LogFolder implement
func (w WebApp) LogFolder() string {
	return filepath.Join(w.BaseFolder(), "log")
}

// HttpFolder implement
func (w WebApp) HttpFolder() string {
	return filepath.Join(w.BaseFolder(), "http")
}

// ConsoleFolder implement
func (w WebApp) ConsoleFolder() string {
	return filepath.Join(w.BaseFolder(), "console")
}

// StorageFolder implement
func (w WebApp) StorageFolder() string {
	return filepath.Join(w.BaseFolder(), "storage")
}

// MiddlewareFolder implment
func (w WebApp) MiddlewareFolder() string {
	return filepath.Join(w.BaseFolder(), "middleware")
}

// ProviderFolder implement
func (w WebApp) ProviderFolder() string {
	return filepath.Join(w.BaseFolder(), "provider")
}


// CommandFolder implment
func (w WebApp) CommandFolder() string {
	return filepath.Join(w.BaseFolder(), "command")
}

// RuntimeFolder implment
func (w WebApp) RuntimeFolder() string {
	return filepath.Join(w.BaseFolder(), "runtime")
}

// TestFolder implement
func (w WebApp) TestFolder() string {
	return filepath.Join(w.BaseFolder(), "test")
}

// NewWebApp
func NewWebApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// Two parameters first is container second is baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &WebApp{container: container, baseFolder: baseFolder}, nil
}







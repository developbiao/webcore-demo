package contract

// AppKey identify
const AppKey = "web:app"

// App define app interface
type App interface {
	// Version
	Version() string
	// BaseFolder project base path
	BaseFolder() string
	// ConfigFoler Config file path
	ConfigFolder() string
	// LogFolder log file path
	LogFolder() string
	// ProviderFolder provider folder
	ProviderFolder() string
	// MiddlewareFolder middleware folder
	MiddlewareFolder() string
	// CommandFolder command folder
	CommandFolder() string
	// RuntimeFolder runtime folder
	RuntimeFolder() string
	// TestFolder test folder
	TestFolder() string
}

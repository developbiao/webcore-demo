package framework

// NewInstance create all container service
type NewInstance func(...interface{}) (interface{}, error)

// ServiceProvider define Service providder
type ServiceProvider interface {
	// Register register instance on container
	Register(Container) NewInstance

	// Boot return error service instance fail
	Boot(Container) error

	// IsDefer false don't need delay initialization
	// true need delay initialization
	IsDefer() bool

	// Params define new instance paramters
	Params(Container) []interface{}

	// Name represent service identify
	Name() string
}

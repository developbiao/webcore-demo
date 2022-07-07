package framework

// NewInstance create all container service
type NewInstance func(...interface{}) (interface{}, error)

// ServiceProvider define Service providder
type ServiceProvider interface {
	// Regisgter resiter instance on container
	Register(Container) NewInstance

	// Booth return error service instance fiaure
	Boot(Container) error

	// IsDefer false don't need delay initlization
	// true need delay initilization
	IsDefer() bool

	// Params define new instance paramters
	Params(Container) []interface{}

	// Name represent service identify
	Name() string
}

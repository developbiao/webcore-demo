package framework

// Represent prefx gorup
type IGroup interface {
	// Implement HttpMethod
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	// Implement Group embedding
	Group(string) IGroup
}

// Group struct
type Group struct {
	core   *Core  // point core structure
	parent *Group // if exists point preve group
	prefix string // group prefix
}

// NewGroup
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

// Get current gorup absolute prefix
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

// Get method
func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Get(uri, handler)
}

// Post
func (g *Group) Post(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Post(uri, handler)
}

// Put
func (g *Group) Put(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Put(uri, handler)
}

// Delete
func (g *Group) Delete(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Delete(uri, handler)
}

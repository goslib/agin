package agin

func NewRouter(key, name, path string, modules Routers, handlers Routes, description string) *Router {
	r := &Router{df(key, name, path, description), nil, modules, handlers}
	for i := 0; i < len(modules); i++ {
		modules[i].Parent = r
	}
	for i := 0; i < len(handlers); i++ {
		handlers[i].Parent = r
	}
	return r
}

func NewGroupedRouter(key, name, path, description string, sub ...*Router) *Router {
	return NewRouter(key, name, path, sub, nil, description)
}

func NewEndingRouter(key, name, path, description string, sub ...*Route) *Router {
	return NewRouter(key, name, path, nil, sub, description)
}

var _ IRestEntry = (*Router)(nil)

// [ Session Router [ Cooperation > App > Module ]  >>  Resource Route >> Route Handler ]
type Routers = []*Router
type Router struct {
	*EmbeddedEntry

	Parent *Router

	SubRouters []*Router

	Routes []*Route
}

// [ Use / Mount / Initialize ]
func (m *Router) Use(r *RouterGroup) {
	if m.Path != "" {
		r = r.Group(m.Path)
	}

	for i := 0; i < len(m.SubRouters); i++ {
		m.SubRouters[i].Use(r)
	}

	for i := 0; i < len(m.Routes); i++ {
		m.Routes[i].Handle(r)
	}
}

func (m *Router) GetPath() string {
	if m.Parent == nil {
		return m.Path
	}
	return m.Parent.GetPath() + m.Path
}

func (m *Router) GetTag() string {
	if m.Parent == nil {
		return m.Name
	}
	return m.Parent.GetTag() + conf.TaggingSeparator + m.Name
}

package agin

import (
	"fmt"
	"net/http"
)

type funcRouteHandler = func(env *Route, ctx *Context) interface{}
type funcGetHandler = func(env *Route) func(ctx *Context)
type TypeLogicalUnit = funcGetHandler

func NewRoute(name, description, method, path string, unit TypeLogicalUnit) *Route {
	return &Route{df("", name, path, description), method, nil, unit, ""}
}

func NewGetRoute(name, description, path string, unit TypeLogicalUnit) *Route {
	return NewRoute(name, description, http.MethodGet, path, unit)
}

func NewPostRoute(name, description, path string, unit TypeLogicalUnit) *Route {
	return NewRoute(name, description, http.MethodPost, path, unit)
}

func NewPatchRoute(name, description, path string, unit TypeLogicalUnit) *Route {
	return NewRoute(name, description, http.MethodPatch, path, unit)
}

func NewPutRoute(name, description, path string, unit TypeLogicalUnit) *Route {
	return NewRoute(name, description, http.MethodPut, path, unit)
}

func NewDeleteRoute(name, description, path string, unit TypeLogicalUnit) *Route {
	return NewRoute(name, description, http.MethodDelete, path, unit)
}

var _ IRestEntry = (*Route)(nil)

type Routes = []*Route
type Route struct {
	*EmbeddedEntry
	Method string

	Parent *Router
	//Handler func(env *Route, ctx *Context) interface{}
	//LogicalUnit ILogicalUnit
	//NewProcessor func(ctx *Context) IProcessor
	GetHandler func(env *Route) func(ctx *Context)

	HandlerName string
}

func (m *Route) Handle(r *RouterGroup) {
	//r.Handle(m.Method, m.Path, NewHandler(m).GetHandler())
	//r.Handle(entry.Method, entry.Path, entry.LogicalUnit.GetHandler())
	//r.Handle(m.Method, m.Path, m.LogicalUnit.Handler)
	//r.Handle(m.Method, m.Path, func(ctx *Context) {
	//	m.Handler(m, ctx)
	//})
	r.Handle(m.Method, m.Path, m.GetHandler(m))
	if conf.Debug {
		fmt.Println("[REST/GIN/"+m.Method+"]\t   ->> ", m.GetPath(), " <<- 【"+m.GetTag()+"】\t", m.Description, "\t#"+m.HandlerName)
	}
}

func (m *Route) GetPath() string {
	if m.Parent == nil {
		return m.Path
	}
	return m.Parent.GetPath() + m.Path
}

func (m *Route) GetTag() string {
	if m.Parent == nil {
		return m.Name
	}
	return m.Parent.GetTag() + conf.TaggingSeparator + m.Name
}

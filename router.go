package crux

import "strings"

const (
	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodDelete  = "DELETE"
	MethodPatch   = "PATCH"
	MethodHead    = "HEAD"
	MethodOptions = "OPTIONS"
)

type Group struct {
	prefix string
	app    *App
}

func (g *Group) Group(prefix string) *Group {
	fullPrefix := g.prefix + "/" + strings.Trim(prefix, "/")
	return &Group{
		prefix: fullPrefix,
		app:    g.app,
	}
}

func (a *App) Group(prefix string) *Group {
	return &Group{
		prefix: strings.Trim(prefix, "/"),
		app:    a,
	}
}

func (g *Group) GET(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.GET(fullPath, handler)
}

func (g *Group) POST(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.POST(fullPath, handler)
}

func (g *Group) PUT(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.PUT(fullPath, handler)
}

func (g *Group) DELETE(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.DELETE(fullPath, handler)
}

func (g *Group) PATCH(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.PATCH(fullPath, handler)
}

func (g *Group) HEAD(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.HEAD(fullPath, handler)
}

func (g *Group) OPTIONS(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.OPTIONS(fullPath, handler)
}

func (a *App) handle(method, path string, handler interface{}) {
	var h Handler

	switch v := handler.(type) {
	case string:
		h = func(c *Context) {
			c.Send(v)
		}
	case Handler:
		h = v
	case func(*Context):
		h = v
	default:
		panic("handler must be either a string or Handler")
	}

	parts := strings.Split(strings.Trim(path, "/"), "/")
	a.routes = append(a.routes, Route{
		path:    path,
		handler: h,
		method:  method,
		parts:   parts,
	})
}

func (a *App) GET(path string, handler interface{}) {
	a.handle(MethodGet, path, handler)
}

func (a *App) POST(path string, handler interface{}) {
	a.handle(MethodPost, path, handler)
}

func (a *App) PUT(path string, handler interface{}) {
	a.handle(MethodPut, path, handler)
}

func (a *App) DELETE(path string, handler interface{}) {
	a.handle(MethodDelete, path, handler)
}

func (a *App) PATCH(path string, handler interface{}) {
	a.handle(MethodPatch, path, handler)
}

func (a *App) HEAD(path string, handler interface{}) {
	a.handle(MethodHead, path, handler)
}

func (a *App) OPTIONS(path string, handler interface{}) {
	a.handle(MethodOptions, path, handler)
}

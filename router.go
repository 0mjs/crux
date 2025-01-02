package crux

import (
	"strings"
)

const (
	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodHead    = "HEAD"
	MethodOptions = "OPTIONS"
	MethodConnect = "CONNECT"
	MethodTrace   = "TRACE"
)

type RouteNode struct {
	path     string
	part     string
	children []*RouteNode
	handler  RouteHandler
	isParam  bool
	isWild   bool
}

type Route struct {
	path    string
	handler RouteHandler
	method  string
	parts   []string
}

type Middleware func(c *Context)

type Router struct {
	routes []*Route
	router *RouteNode
}

func (r *Router) Add(method, path string, handlers ...interface{}) {
	var routeHandlers []RouteHandler

	for _, handler := range handlers {
		var rh RouteHandler

		switch v := handler.(type) {
		case string:
			rh = func(c *Context) {
				c.Send(v)
			}
		case RouteHandler:
			rh = v
		case func(*Context):
			rh = v
		case Middleware:
			rh = RouteHandler(v)
		default:
			panic("handler must be either a string, RouteHandler, or Middleware")
		}

		routeHandlers = append(routeHandlers, rh)
	}

	var mainHandler RouteHandler
	if len(routeHandlers) > 0 {
		mainHandler = chain(routeHandlers)
	}

	path = r.normalizePath(path)

	if r.router == nil {
		r.router = &RouteNode{}
	}

	parts := strings.Split(strings.Trim(path, "/"), "/")
	current := r.router

	for i, part := range parts {
		isParam := false
		isWild := false

		if len(part) > 0 {
			if part[0] == ':' {
				isParam = true
				part = part[1:]
			} else if part[0] == '*' {
				isWild = true
				part = part[1:]
			}
		}

		child := current.findChild(part, isParam, isWild)
		if child == nil {
			child = &RouteNode{
				part:    part,
				isParam: isParam,
				isWild:  isWild,
			}
			current.children = append(current.children, child)
		}

		if i == len(parts)-1 {
			child.handler = mainHandler
			child.path = path
		}

		current = child
	}

	r.routes = append(r.routes, &Route{
		path:    path,
		handler: mainHandler,
		method:  method,
		parts:   parts,
	})
}

func chain(handlers []RouteHandler) RouteHandler {
	return func(c *Context) {
		for _, handler := range handlers {
			handler(c)
			if c.written {
				return
			}
		}
	}
}

func (n *RouteNode) findChild(part string, isParam, isWild bool) *RouteNode {
	for _, child := range n.children {
		if child.part == part && child.isParam == isParam && child.isWild == isWild {
			return child
		}
	}
	return nil
}

func (r *Router) normalizePath(path string) string {
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		return "/" + path
	}
	return path
}

func (r *Router) Find(path string) (RouteHandler, map[string]string) {
	params := make(map[string]string)
	node := r.router.find(strings.Split(strings.Trim(path, "/"), "/"), params)
	if node != nil {
		return node.handler, params
	}
	return nil, nil
}

func (n *RouteNode) find(parts []string, params map[string]string) *RouteNode {
	if len(parts) == 0 {
		return n
	}

	part := parts[0]
	parts = parts[1:]

	for _, child := range n.children {
		if child.isWild {
			params["*"] = strings.Join(append([]string{part}, parts...), "/")
			return child
		}

		if child.isParam {
			params[child.part] = part
			if matchChild := child.find(parts, params); matchChild != nil {
				return matchChild
			}
		} else if child.part == part {
			if matchChild := child.find(parts, params); matchChild != nil {
				return matchChild
			}
		}
	}

	return nil
}

func (a *App) GET(path string, handlers ...interface{}) {
	a.router.Add(MethodGet, path, handlers...)
}

func (a *App) POST(path string, handlers ...interface{}) {
	a.router.Add(MethodPost, path, handlers...)
}

func (a *App) PUT(path string, handlers ...interface{}) {
	a.router.Add(MethodPut, path, handlers...)
}

func (a *App) DELETE(path string, handlers ...interface{}) {
	a.router.Add(MethodDelete, path, handlers...)
}

func (a *App) PATCH(path string, handlers ...interface{}) {
	a.router.Add(MethodPatch, path, handlers...)
}

func (a *App) HEAD(path string, handlers ...interface{}) {
	a.router.Add(MethodHead, path, handlers...)
}

func (a *App) OPTIONS(path string, handlers ...interface{}) {
	a.router.Add(MethodOptions, path, handlers...)
}

func (a *App) CONNECT(path string, handlers ...interface{}) {
	a.router.Add(MethodConnect, path, handlers...)
}

func (a *App) TRACE(path string, handlers ...interface{}) {
	a.router.Add(MethodTrace, path, handlers...)
}

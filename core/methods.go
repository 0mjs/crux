package crux

import "strings"

func (a *App) Get(path string, handler interface{}) {
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
		method:  "GET",
		parts:   parts,
	})
}

func (a *App) Post(path string, handler Handler) {
	a.routes = append(a.routes, Route{
		path:    path,
		handler: handler,
		method:  "POST",
	})
}

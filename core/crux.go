package crux

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Map map[string]interface{}

type Context struct {
	Response    http.ResponseWriter
	Request     *http.Request
	PathParams  map[string]string
	QueryParams url.Values
}

type Handler func(ctx *Context)

type Route struct {
	path    string
	handler Handler
	method  string
	parts   []string
}

type App struct {
	routes []Route
}

func New() *App {
	return &App{
		routes: make([]Route, 0),
	}
}

func (c *Context) JSON(data interface{}) error {
	c.Response.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(c.Response).Encode(data)
}

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

func (c *Context) Send(data interface{}) error {
	switch v := data.(type) {
	case string:
		c.Response.Header().Set("Content-Type", "text/plain")
		_, err := c.Response.Write([]byte(v))
		return err
	case []byte:
		c.Response.Header().Set("Content-Type", "application/octet-stream")
		_, err := c.Response.Write(v)
		return err
	default:
		c.Response.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(c.Response).Encode(data)
	}
}

func (c *Context) Param(name string) string {
	return c.PathParams[name]
}

func (c *Context) Query(name string) string {
	return c.QueryParams.Get(name)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		Response:    w,
		Request:     r,
		PathParams:  make(map[string]string),
		QueryParams: r.URL.Query(),
	}

	path := r.URL.Path

	for _, route := range a.routes {
		if matchPath(path, route, ctx) && route.method == r.Method {
			route.handler(ctx)
			return
		}
	}
	http.NotFound(w, r)
}

func matchPath(path string, route Route, ctx *Context) bool {
	pathParts := strings.Split(strings.Trim(path, "/"), "/")
	routeParts := strings.Split(strings.Trim(route.path, "/"), "/")

	if len(pathParts) != len(routeParts) {
		return false
	}

	for i, part := range routeParts {
		if strings.HasPrefix(part, ":") {
			paramName := strings.TrimPrefix(part, ":")
			ctx.PathParams[paramName] = pathParts[i]
			continue
		}
		if part != pathParts[i] {
			return false
		}
	}
	return true
}

func (a *App) Listen(port int) error {
	fmt.Printf("Server starting on port %d...\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), a)
}

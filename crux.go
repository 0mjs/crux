package crux

import (
	"fmt"
	"net/http"
	"strings"
)

type App struct {
	routes []Route
}

type Handler func(ctx *Context)

type Route struct {
	path    string
	handler Handler
	method  string
	parts   []string
}

type Map map[string]interface{}

func New() *App {
	return &App{
		routes: make([]Route, 0),
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		Response:    w,
		Request:     r,
		PathParams:  make(map[string]string),
		QueryParams: r.URL.Query(),
		Method:      r.Method,
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

func (a *App) Listen(port ...int) error {
	defaultPort := 8080
	if len(port) > 0 && port[0] != 0 {
		defaultPort = port[0]
	} else {
		fmt.Printf("No port provided, using default port %d\n", defaultPort)
	}
	fmt.Printf("Server starting on port %d...\n", defaultPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), a)
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
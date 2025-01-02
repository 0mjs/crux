package crux

import (
	"fmt"
	"net/http"
	"strings"
)

type App struct {
	router     *Router
	middleware []Middleware
}

type RouteHandler func(ctx *Context)

type Map map[string]interface{}

func New() *App {
	return &App{
		router:     &Router{},
		middleware: make([]Middleware, 0),
	}
}

func (a *App) Use(middleware ...Middleware) {
	a.middleware = append(a.middleware, middleware...)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)

	ctx.setHandlers(a.middleware)

	if len(a.middleware) > 0 {
		ctx.Next()
		if ctx.written {
			return
		}
	}

	path := r.URL.Path

	for _, m := range a.middleware {
		m(ctx)
		if ctx.written {
			return
		}
	}

	for _, route := range a.router.routes {
		if matchPath(path, *route, ctx) && route.method == r.Method {
			route.handler(ctx)
			return
		}
	}
	http.NotFound(w, r)
}

func (a *App) Listen(port ...int) error {
	fmt.Printf("App Router: %+v\n", a.router)

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

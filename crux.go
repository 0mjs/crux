package crux

import (
	"fmt"
	"net/http"
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

	handler, params := a.router.Find(r.URL.Path)
	if handler != nil && params != nil {
		ctx.PathParams = params
		handler(ctx)
		return
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

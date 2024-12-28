package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerFunc is a custom type for route handlers
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Route represents a single route with its path and handler
type Route struct {
	path    string
	handler HandlerFunc
	method  string
}

// App is the main application struct
type App struct {
	routes []Route
}

// Crux creates a new instance of the application
func Crux() *App {
	return &App{
		routes: make([]Route, 0),
	}
}

// Get registers a new GET route
func (a *App) Get(path string, handler HandlerFunc) {
	a.routes = append(a.routes, Route{
		path:    path,
		handler: handler,
		method:  "GET",
	})
}

// Post registers a new POST route
func (a *App) Post(path string, handler HandlerFunc) {
	a.routes = append(a.routes, Route{
		path:    path,
		handler: handler,
		method:  "POST",
	})
}

// Send is a helper method to send various types of responses
func (a *App) Send(w http.ResponseWriter, data interface{}) error {
	switch v := data.(type) {
	case string:
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte(v))
		return err
	case []byte:
		w.Header().Set("Content-Type", "application/octet-stream")
		_, err := w.Write(v)
		return err
	default:
		// Assume it's JSON-serializable
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(data)
	}
}

// ServeHTTP implements the http.Handler interface
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range a.routes {
		if route.path == r.URL.Path && route.method == r.Method {
			route.handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

// Listen starts the server on the specified port
func (a *App) Listen(port int) error {
	fmt.Printf("Server starting on port %d...\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), a)
}

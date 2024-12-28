package crux

import (
	"net/http"
	"net/url"
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

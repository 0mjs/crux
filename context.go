package crux

import (
	"net/http"
	"net/url"
)

type Context struct {
	Response    http.ResponseWriter
	Request     *http.Request
	PathParams  map[string]string
	QueryParams url.Values
	Method      string
}

func (c *Context) Param(name string) string {
	return c.PathParams[name]
}

func (c *Context) Query(name string) string {
	return c.QueryParams.Get(name)
}

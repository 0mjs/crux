package crux

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Context struct {
	Response    http.ResponseWriter
	Request     *http.Request
	PathParams  map[string]string
	QueryParams url.Values
	Method      string
	written     bool
	handlers    []Middleware
	index       int
	Store       map[string]interface{}
	status      int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response:    w,
		Request:     r,
		PathParams:  make(map[string]string),
		QueryParams: r.URL.Query(),
		Method:      r.Method,
		Store:       make(map[string]interface{}),
		status:      http.StatusOK,
	}
}

func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}

func (c *Context) setHandlers(handlers []Middleware) {
	c.handlers = handlers
	c.index = -1
}

func (c *Context) Set(key string, value interface{}) {
	c.Store[key] = value
}

func (c *Context) Get(key string) interface{} {
	return c.Store[key]
}

func (c *Context) Status(code int) *Context {
	c.status = code
	return c
}

func (c *Context) Param(name string) string {
	return c.PathParams[name]
}

func (c *Context) Query(name string) string {
	return c.QueryParams.Get(name)
}

func (c *Context) Body(v interface{}) error {
	if c.Request.Body == nil {
		return errors.New("empty request body")
	}
	defer c.Request.Body.Close()

	if err := json.NewDecoder(c.Request.Body).Decode(v); err != nil {
		return err
	}

	if err := validate.Struct(v); err != nil {
		return err
	}

	return nil
}

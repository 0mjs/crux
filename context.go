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
	response    http.ResponseWriter
	request     *http.Request
	pathParams  map[string]string
	queryParams url.Values
	Method      string
	written     bool
	handlers    []RouteHandler
	index       int
	Store       map[string]interface{}
	status      int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		response:    w,
		request:     r,
		pathParams:  make(map[string]string),
		queryParams: r.URL.Query(),
		Method:      r.Method,
		Store:       make(map[string]interface{}),
		status:      http.StatusOK,
	}
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
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
	return c.pathParams[name]
}

func (c *Context) Query(name string) string {
	return c.queryParams.Get(name)
}

func (c *Context) Body(v interface{}) error {
	if c.request.Body == nil {
		return errors.New("empty request body")
	}
	defer c.request.Body.Close()

	if err := json.NewDecoder(c.request.Body).Decode(v); err != nil {
		return err
	}

	if err := validate.Struct(v); err != nil {
		return err
	}

	return nil
}

package crux

import (
	"encoding/json"
	"net/http"
)

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

func (c *Context) JSON(data interface{}) error {
	c.written = true
	c.Response.Header().Set("Content-Type", "application/json")

	if c.status == 0 {
		c.status = http.StatusOK
	}

	c.Response.WriteHeader(c.status)
	return json.NewEncoder(c.Response).Encode(data)
}

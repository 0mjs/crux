package crux

import (
	"encoding/json"
	"net/http"
)

func (c *Context) Send(data interface{}) error {
	switch v := data.(type) {
	case string:
		c.response.Header().Set("Content-Type", "text/plain")
		_, err := c.response.Write([]byte(v))
		return err
	case []byte:
		c.response.Header().Set("Content-Type", "application/octet-stream")
		_, err := c.response.Write(v)
		return err
	default:
		c.response.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(c.response).Encode(data)
	}
}

func (c *Context) JSON(data interface{}) error {
	c.written = true
	c.response.Header().Set("Content-Type", "application/json")

	if c.status == 0 {
		c.status = http.StatusOK
	}

	c.response.WriteHeader(c.status)
	return json.NewEncoder(c.response).Encode(data)
}

package main

import (
	"fmt"

	"github.com/0mjs/crux"
)

func main() {
	app := crux.New()

	hello := func(c *crux.Context) {
		c.Send("Hello, to the World!")
	}

	// Implicit string response
	app.GET("/", "Hello, World!")

	app.GET("/hello", hello)

	// Explicit string response
	app.GET("/alt", func(c *crux.Context) {
		c.Send("Hello, World!")
	})

	// Inferred JSON response
	app.GET("/json1", func(c *crux.Context) {
		c.Send(crux.Map{
			"message": "Hello, JSON!",
		})
	})

	// JSON response
	app.GET("/json2", func(c *crux.Context) {
		c.JSON(crux.Map{
			"message": "Hello, JSON!",
		})
	})

	// Path parameters
	app.GET("/users/:id", func(c *crux.Context) {
		c.JSON(crux.Map{
			"message": fmt.Sprintf("the user id is %s", c.Param("id")),
		})
	})

	// Nested path parameters
	app.GET("/users/:userID/posts/:postID", func(c *crux.Context) {
		c.JSON(crux.Map{
			"user": c.Param("userID"),
			"post": c.Param("postID"),
		})
	})

	// Query parameters
	app.GET("/search", func(c *crux.Context) {
		c.JSON(crux.Map{
			"search": c.Query("q"),
			"limit":  c.Query("limit"),
		})
	})

	// Route grouping
	api := app.Group("/api")

	api.GET("/usernames", func(c *crux.Context) {
		c.JSON(crux.Map{
			"usernames": []string{"martin", "jason"},
		})
	})

	// Nested route grouping
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")

	v1Group := v1.Group("/users")
	v1Group.GET("/", func(c *crux.Context) {
		c.JSON(crux.Map{
			"version": "v1",
			"users":   []string{"martin", "jason"},
		})
	})

	v2Group := v2.Group("/users")
	v2Group.GET("/", func(c *crux.Context) {
		c.JSON(crux.Map{
			"version": "v2",
			"users":   []string{"martin", "jason", "peter"},
		})
	})

	// Method routing

	app.GET("/get-method", func(c *crux.Context) {
		c.JSON(crux.Map{
			"method": c.Method,
		})
	})

	app.POST("/post-method", func(c *crux.Context) {
		c.JSON(crux.Map{
			"method": c.Method,
		})
	})

	app.PUT("/put-method", func(c *crux.Context) {
		c.JSON(crux.Map{
			"method": c.Method,
		})
	})

	app.DELETE("/delete-method", func(c *crux.Context) {
		c.JSON(crux.Map{
			"method": c.Method,
		})
	})

	app.PATCH("/patch-method", func(c *crux.Context) {
		c.JSON(crux.Map{
			"method": c.Method,
		})
	})

	app.HEAD("/head-method", func(c *crux.Context) {
		c.JSON(crux.Map{
			"method": c.Method,
		})
	})

	app.OPTIONS("/options-method", func(c *crux.Context) {
		c.JSON(crux.Map{
			"method": c.Method,
		})
	})

	app.Listen()
}

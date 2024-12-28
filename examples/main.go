package main

import (
	"fmt"

	crux "github.com/0mjs/crux/core"
)

func main() {
	app := crux.New()

	// Test string response, with IMPLICIT .Send() (assumes string to be used as response)
	app.Get("/", "Hello, World!")

	// Test string response, with EXPLICIT .Send()
	app.Get("/alt", func(c *crux.Context) {
		c.Send("Hello, World!")
	})

	// Test JSON response, with .Send()
	app.Get("/json1", func(c *crux.Context) {
		c.Send(crux.Map{
			"message": "Hello, JSON!",
		})
	})

	// Test JSON response, with .JSON()
	app.Get("/json2", func(c *crux.Context) {
		c.JSON(crux.Map{
			"message": "Hello, JSON!",
		})
	})

	// Path parameters
	app.Get("/users/:id", func(c *crux.Context) {
		userID := c.Param("id")
		c.JSON(crux.Map{
			"message": fmt.Sprintf("Got user %s", userID),
		})
	})

	// Multiple path parameters
	app.Get("/users/:userID/posts/:postID", func(c *crux.Context) {
		userID := c.Param("userID")
		postID := c.Param("postID")
		c.JSON(crux.Map{
			"user": userID,
			"post": postID,
		})
	})

	// Query parameters
	app.Get("/search", func(c *crux.Context) {
		query := c.Query("q")
		limit := c.Query("limit")
		c.JSON(crux.Map{
			"search": query,
			"limit":  limit,
		})
	})

	app.Listen()
}

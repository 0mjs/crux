package main

import (
	"fmt"
	"slices"

	"github.com/0mjs/crux"
)

func Authenticate() crux.Middleware {
	return func(c *crux.Context) {
		c.Set("authenticated", true)
		c.Next()
	}
}

func Authorize(permission string) crux.Middleware {
	return func(c *crux.Context) {
		validPermissions := []string{"some-permission", "another-permission"}

		if !slices.Contains(validPermissions, permission) {
			c.Status(403).JSON(crux.Map{"error": "Unauthorized"})
			return
		}

		c.Set("authorized", permission)
		c.Next()
	}
}

func exampleMiddleware() crux.Middleware {
	fmt.Println("Middleware 'exampleMiddleware' initialized")
	return func(c *crux.Context) {
		fmt.Println("Request received:", c.Request.URL.Path)
		c.Next()
	}
}

func main() {
	app := crux.New()

	app.Use(exampleMiddleware())

	app.GET(
		"/",
		Authenticate(),
		Authorize("some-permission"),
		func(c *crux.Context) {
			authenticated := c.Get("authenticated")
			authorized := c.Get("authorized")
			c.JSON(crux.Map{
				"message":       "Hello, World!",
				"authenticated": authenticated,
				"authorized":    authorized,
			})
		},
	)

	app.Listen(8080)
}

package main

import (
	"fmt"
	"slices"

	"github.com/0mjs/crux"
)

type Middleware func(c *crux.Context)

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

func Dink() crux.Middleware {
	return func(c *crux.Context) {
		fmt.Println(c.Store)
	}
}

func main() {
	app := crux.New()

	app.GET(
		"/",
		Authenticate(),
		Authorize("some-permission"),
		Dink(),
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

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

func CustomHTML(c *crux.Context) {
	html := `<!DOCTYPE html><html lang="en"><head> <meta charset="UTF-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <title>Hello World</title> <style> body { margin: 0; height: 100vh; display: flex; align-items: center; justify-content: center; background: linear-gradient(135deg, #6366f1, #a855f7); font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, sans-serif; } .container { background: rgba(255, 255, 255, 0.95); padding: 2rem 3rem; border-radius: 1rem; box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04); text-align: center; } h1 { color: #1f2937; margin: 0; font-size: 2.5rem; font-weight: 700; } p { color: #4b5563; margin-top: 1rem; font-size: 1.1rem; } </style></head><body> <div class="container"> <h1>Hello World!</h1> <p>Welcome to your styled endpoint</p> </div></body></html>`

	c.HTML(html)
}

func exampleMiddleware() crux.Middleware {
	fmt.Println("Middleware 'exampleMiddleware' initialized via app.Use()")
	return func(c *crux.Context) {
		fmt.Println("Request received:", c.Request.URL.Path)
		c.Next()
	}
}

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

	// Route-level middleware
	app.Use(exampleMiddleware())

	// Chained middleware
	app.GET(
		"/permissions",
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

	// Custom HTML
	app.GET("/html", CustomHTML)

	// Static file serving
	app.GET("/static", func(c *crux.Context) {
		c.Static("eg/static/index.html")
	})

	app.Listen()
}

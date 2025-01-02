package main

import (
	"github.com/0mjs/crux"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=130"`
}

func main() {
	app := crux.New()

	app.POST("/", func(c *crux.Context) {
		var user User
		if err := c.Body(&user); err != nil {
			c.JSON(crux.Map{
				"error": err.Error(),
			})
			return
		}
		c.JSON(user)
	})

	app.Listen(8080)
}

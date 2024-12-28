package main

import (
	"net/http"

	"github.com/0mjs/crux/pkg/crux"
)

func main() {
	app := crux.Crux()

	// Simple string response
	app.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		app.Send(w, "Hello, World!")
	})

	// JSON response
	app.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{"message": "Hello, JSON!"}
		app.Send(w, data)
	})

	// Start the server
	app.Listen(3000)
}

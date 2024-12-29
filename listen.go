package crux

import (
	"fmt"
	"net/http"
)

func (a *App) Listen(port ...int) error {
	defaultPort := 8080
	if len(port) > 0 && port[0] != 0 {
		defaultPort = port[0]
	} else {
		fmt.Printf("No port provided, using default port %d\n", defaultPort)
	}
	fmt.Printf("Server starting on port %d...\n", defaultPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), a)
}

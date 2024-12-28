package crux

import (
	"fmt"
	"net/http"
)

func (a *App) Listen(port ...int) error {
	serverPort := 8080
	if len(port) > 0 && port[0] != 0 {
		serverPort = port[0]
	} else {
		fmt.Printf("No port provided, using default port %d\n", serverPort)
	}
	fmt.Printf("Server starting on port %d...\n", serverPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", serverPort), a)
}

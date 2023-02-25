package main

import (
	"fmt"
	"goAnsible/pkg/handlers"
	"net/http"
)

const portNumber = ":8089"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

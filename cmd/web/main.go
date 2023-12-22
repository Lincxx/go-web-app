package main

import (
	"fmt"
	"github.com/Lincxx/go-web-app/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port: ", portNumber)
	//web server
	_ = http.ListenAndServe(portNumber, nil)
}

package main

import (
	"fmt"
	"github.com/Lincxx/go-web-app/pkg/config"
	"github.com/Lincxx/go-web-app/pkg/handlers"
	"github.com/Lincxx/go-web-app/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port!: ", portNumber)
	//web server
	_ = http.ListenAndServe(portNumber, nil)
}

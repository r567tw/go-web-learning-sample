package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/r567tw/go-api-service/pkg/config"
	"github.com/r567tw/go-api-service/pkg/handler"
	"github.com/r567tw/go-api-service/pkg/render"
)

const portNumber = ":8081"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
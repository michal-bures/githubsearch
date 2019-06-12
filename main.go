package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"githubsearch/controller"
	"githubsearch/searcher"
	"log"
)
import "net/http"

func main() {
	loadEnvVariablesFromEnvFile()
	searchEngine := searcher.NewSearcher()
	controller := controller.NewController(searchEngine)
	startHttpServer(controller)
}

func startHttpServer(controller *controller.Controller) {
	http.HandleFunc("/", controller.IndexPageHandler)

	//	fs := http.FileServer(http.Dir("static/"))
	//	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Print("Starting server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}

func loadEnvVariablesFromEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Println("File .env not found, running with no extra ENV variables")
	}
}

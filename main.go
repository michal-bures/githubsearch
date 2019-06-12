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
	searcher := searcher.NewSearcher()
	controller := controller.NewController(searcher)
	startHttpServer(controller)
}

func startHttpServer(controller *controller.Controller) {
	http.HandleFunc("/", controller.IndexPageHandler)

	fmt.Print("Listening on port 8080")
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

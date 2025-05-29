package main

import (
	"log"
	"net/http"
	"os"

	"github.com/datslim/quote-api/internal/controller"
)

func main() {
	logger := log.New(os.Stdout, "QUOTE-APP: ", log.LstdFlags|log.Lshortfile)

	router := controller.NewRouter()

	logger.Printf("Сервер запущен на порту 8080.\n")
	logger.Printf("http://localhost:8080/quotes")
	logger.Fatal(http.ListenAndServe(":8080", router))
}

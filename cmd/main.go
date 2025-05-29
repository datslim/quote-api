package main

import (
	"log"
	"net/http"
	"os"

	"github.com/datslim/quote-api/internal/handlers"
	"github.com/datslim/quote-api/internal/storage"
	"github.com/gorilla/mux"
)

func main() {
	store := storage.NewMemoryStorage()
	logger := log.New(os.Stdout, "QUOTE-API: ", log.LstdFlags|log.Lshortfile)
	quoteHandler := handlers.NewQuoteHandler(store, logger)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/quotes", quoteHandler.PostQuote).Methods("POST")
	router.HandleFunc("/quotes", quoteHandler.GetAllQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", quoteHandler.GetRandomQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", quoteHandler.DeleteQuote).Methods("DELETE")
	logger.Printf("Сервер запущен на порту 8080.\n")
	logger.Printf("http://localhost:8080/quotes")
	logger.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"log"
	"net/http"

	"github.com/datslim/quote-api/internal/handlers"
	"github.com/datslim/quote-api/internal/storage"
	"github.com/gorilla/mux"
)

func main() {
	store := storage.NewMemoryStorage()

	quoteHandler := handlers.NewQuoteHandler(store)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/quotes", quoteHandler.PostQuote).Methods("POST")
	router.HandleFunc("/quotes", quoteHandler.GetAllQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", quoteHandler.GetRandomQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", quoteHandler.DeleteQuote).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

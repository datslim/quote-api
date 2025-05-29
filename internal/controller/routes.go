package controller

import (
	"log"
	"os"

	"github.com/datslim/quote-api/internal/handlers"
	"github.com/datslim/quote-api/internal/storage"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	logger := log.New(os.Stdout, "HTTP: ", log.LstdFlags|log.Lshortfile)
	store := storage.NewMemoryStorage()
	quoteHandler := handlers.NewQuoteHandler(store, logger)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/quotes", quoteHandler.PostQuote).Methods("POST")
	router.HandleFunc("/quotes", quoteHandler.GetAllQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", quoteHandler.GetRandomQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", quoteHandler.DeleteQuote).Methods("DELETE")

	return router
}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/datslim/quote-api/internal/models"
	"github.com/datslim/quote-api/internal/storage"
	"github.com/gorilla/mux"
)

type QuoteHandler struct {
	store *storage.MemoryStorage
}

func NewQuoteHandler(store *storage.MemoryStorage) *QuoteHandler {
	return &QuoteHandler{store: store}
}

func (h *QuoteHandler) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	var quotes []models.Quote
	if author == "" {
		quotes = h.store.GetAll()
	} else {
		quotes = h.store.GetByAuthor(author)
	}
	json.NewEncoder(w).Encode(quotes)
}

func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.store.GetRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(quote)
}

func (h *QuoteHandler) PostQuote(w http.ResponseWriter, r *http.Request) {
	var quote models.Quote

	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newQuote := h.store.Add(quote)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newQuote)
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if !h.store.Delete(id) {
		http.Error(w, "Quote not found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusNoContent)
}

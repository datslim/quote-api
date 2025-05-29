package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/datslim/quote-api/internal/models"
	"github.com/datslim/quote-api/internal/storage"
	"github.com/gorilla/mux"
)

type QuoteHandler struct {
	store  *storage.MemoryStorage
	logger *log.Logger
}

func NewQuoteHandler(store *storage.MemoryStorage, logger *log.Logger) *QuoteHandler {
	return &QuoteHandler{
		store:  store,
		logger: logger,
	}
}

func (h *QuoteHandler) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	author := r.URL.Query().Get("author")

	var quotes []models.Quote
	if author == "" {
		h.logger.Printf("Показ всех цитат\n")
		quotes = h.store.GetAll()
	} else {
		h.logger.Printf("Показ всех цитат автора (%s)\n", author)
		quotes = h.store.GetByAuthor(author)
	}
	json.NewEncoder(w).Encode(quotes)

	h.logger.Printf("Показано %d цитат. Время выполнения: %v\n", len(quotes), time.Since(start))
}

func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	quote, err := h.store.GetRandom()
	h.logger.Printf("Запрос случайной цитаты\n")
	if err != nil {
		h.logger.Printf("Нет доступных цитат\n")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(quote)

	h.logger.Printf("Возвращена случайная цитата с ID: %d. Время выполнения: %v", quote.ID, time.Since(start))
}

func (h *QuoteHandler) PostQuote(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var quote models.Quote

	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.logger.Printf("Создание цитаты: автор '%s', текст '%s'", quote.Author, quote.Quote)
	newQuote := h.store.Add(quote)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newQuote)
	h.logger.Printf("Создана цитата с ID: %d. Время выполнения: %v", newQuote.ID, time.Since(start))
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	h.logger.Printf("Попытка удаления цитаты с ID: %d\n", id)
	if err != nil {
		h.logger.Printf("Неверный формат ID: %d\n", id)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if !h.store.Delete(id) {
		h.logger.Printf("Цитата с ID: %d не найдена\n", id)
		http.Error(w, "Quote not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	h.logger.Printf("Цитата с ID %d удалена. Время выполнения: %v\n", id, time.Since(start))
}

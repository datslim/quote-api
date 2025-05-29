package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/datslim/quote-api/internal/handlers"
	"github.com/datslim/quote-api/internal/models"
	"github.com/datslim/quote-api/internal/storage"
	"github.com/gorilla/mux"
)

func setupTestHandler() *handlers.QuoteHandler {
	store := storage.NewMemoryStorage()
	store.Add(models.Quote{Author: "Тестовый автор", Quote: "Тестовая цитата"})
	logger := log.New(io.Discard, "", 0)
	return handlers.NewQuoteHandler(store, logger)
}

func TestGetAllQuotes(t *testing.T) {
	handler := setupTestHandler()
	req := httptest.NewRequest("GET", "/quotes", nil)
	w := httptest.NewRecorder()

	handler.GetAllQuotes(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Ожидался код %d, получен %d", http.StatusOK, w.Code)
	}

	var quotes []models.Quote
	if err := json.NewDecoder(w.Body).Decode(&quotes); err != nil {
		t.Fatalf("Ошибка декодирования ответа: %v", err)
	}

	if len(quotes) != 1 {
		t.Errorf("Ожидалось 1 цитата, получено %d", len(quotes))
	}
}

func TestGetRandomQuote(t *testing.T) {
	handler := setupTestHandler()
	req := httptest.NewRequest("GET", "/quotes/random", nil)
	w := httptest.NewRecorder()

	handler.GetRandomQuote(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Ожидался код %d, получен %d", http.StatusOK, w.Code)
	}

	var quote models.Quote
	if err := json.NewDecoder(w.Body).Decode(&quote); err != nil {
		t.Fatalf("Ошибка декодирования ответа: %v", err)
	}

	if quote.ID != 1 {
		t.Errorf("Ожидался ID 1, получен %d", quote.ID)
	}
}

func TestPostQuote(t *testing.T) {
	handler := setupTestHandler()
	newQuote := models.Quote{Author: "Новый автор", Quote: "Новая цитата"}
	body, _ := json.Marshal(newQuote)

	req := httptest.NewRequest("POST", "/quotes", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.PostQuote(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Ожидался код %d, получен %d", http.StatusCreated, w.Code)
	}

	var response models.Quote
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Ошибка декодирования ответа: %v", err)
	}

	if response.ID != 2 {
		t.Errorf("Ожидался ID 2, получен %d", response.ID)
	}
}

func TestDeleteQuote(t *testing.T) {
	handler := setupTestHandler()
	req := httptest.NewRequest("DELETE", "/quotes/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handler.DeleteQuote(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Ожидался код %d, получен %d", http.StatusNoContent, w.Code)
	}
}

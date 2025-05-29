package storage

import (
	"errors"
	"math/rand"
	"slices"

	"github.com/datslim/quote-api/internal/models"
)

type MemoryStorage struct {
	quotes []models.Quote
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		quotes: make([]models.Quote, 0),
	}
}

func (s *MemoryStorage) GetAll() []models.Quote {
	return s.quotes
}

func (s *MemoryStorage) GetRandom() (*models.Quote, error) {
	if len(s.quotes) == 0 {
		return nil, errors.New("no quotes available")
	}
	randomNumber := rand.Intn(len(s.quotes))
	return &s.quotes[randomNumber], nil
}

func (s *MemoryStorage) GetByAuthor(author string) []models.Quote {
	if author == "" {
		return s.quotes
	}

	var filtredQuotes []models.Quote
	for _, quote := range s.quotes {
		if quote.Author == author {
			filtredQuotes = append(filtredQuotes, quote)
		}
	}
	return filtredQuotes
}

func (s *MemoryStorage) Add(quote models.Quote) models.Quote {
	var currentID int = 0
	if len(s.quotes) == 0 {
		currentID = 1
	} else {
		currentID = s.quotes[len(s.quotes)-1].ID + 1
	}
	quote.ID = currentID
	s.quotes = append(s.quotes, quote)
	return quote
}

func (s *MemoryStorage) Delete(id int) bool {
	for i, quote := range s.quotes {
		if quote.ID == id {
			s.quotes = slices.Delete(s.quotes, i, i+1)
			return true
		}
	}
	return false
}

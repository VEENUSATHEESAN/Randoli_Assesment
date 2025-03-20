package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"go-bookstore/models"
	"go-bookstore/storage"
)

// SearchBooks finds books matching the query in title or description
func SearchBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	books, _ := storage.LoadBooks()
	var wg sync.WaitGroup
	results := make(chan models.Book, len(books))

	for _, book := range books {
		wg.Add(1)
		go func(b models.Book) {
			defer wg.Done()
			if strings.Contains(strings.ToLower(b.Title), strings.ToLower(query)) ||
				strings.Contains(strings.ToLower(b.Description), strings.ToLower(query)) {
				results <- b
			}
		}(book)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var matchedBooks []models.Book
	for book := range results {
		matchedBooks = append(matchedBooks, book)
	}

	json.NewEncoder(w).Encode(matchedBooks)
}

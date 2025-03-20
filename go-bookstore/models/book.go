package models

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	BookID          string  `json:"bookId"`
	AuthorID        string  `json:"authorId"`
	PublisherID     string  `json:"publisherId"`
	Title           string  `json:"title"`
	PublicationDate string  `json:"publicationDate"`
	ISBN            string  `json:"isbn"`
	Pages           int     `json:"pages"`
	Genre           string  `json:"genre"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	Quantity        int     `json:"quantity"`
}

func (b *Book) ToJSON() string {
	data, _ := json.Marshal(b)
	return string(data)
}

func (b *Book) Validate() error {
	if b.BookID == "" || b.Title == "" || b.ISBN == "" {
		return fmt.Errorf("bookId, title and isbn are required fields")
	}
	return nil
}

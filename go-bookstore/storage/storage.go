package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"go-bookstore/models"
)

var mutex sync.Mutex

const filePath = "books.json"

func LoadBooks() ([]models.Book, error) {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Book{}, nil
		}
		return nil, err
	}

	var books []models.Book
	err = json.Unmarshal(file, &books)
	return books, err
}

func SaveBooks(books []models.Book) error {
	mutex.Lock()
	defer mutex.Unlock()

	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, 0644)
}

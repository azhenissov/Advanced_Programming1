package Library

import (
	"errors"
	"strconv"
)

type Library struct {
	alexandria map[string]Book
}

func (lib *Library) AddBook(book Book) {
	if lib.alexandria == nil {
		lib.alexandria = make(map[string]Book)
	}
	lib.alexandria[strconv.Itoa(book.ID)] = book
}

func (lib *Library) BorrowBook(ID int) (Book, error) {
	if lib.alexandria == nil {
		return Book{}, errors.New("nothing is in alexandria")
	}

	key := strconv.Itoa(ID)
	book, ok := lib.alexandria[key]
	if !ok {
		return Book{}, errors.New("book not found")
	}
	if book.IsBorrowed {
		return Book{}, errors.New("book is already borrowed")
	}

	book.IsBorrowed = true
	lib.alexandria[key] = book
	return book, nil
}
func (lib *Library) ReturnBook(ID int) error {
	if lib.alexandria == nil {
		return errors.New("nothing is in alexandria")
	}

	key := strconv.Itoa(ID)
	book := lib.alexandria[key]

	if !lib.alexandria[strconv.Itoa(ID)].IsBorrowed {
		return errors.New("book is not borrowed")
	}

	book.IsBorrowed = false
	lib.alexandria[key] = book
	return nil
}

func (lib *Library) ListAvailableBooks() ([]Book, error) {
	if lib.alexandria == nil {
		return []Book{}, errors.New("nothing is in alexandria")
	}
	var books []Book
	for _, book := range lib.alexandria {
		if !book.IsBorrowed {
			books = append(books, book)
		}
	}

	return books, nil
}

package book

import (
	"context"
)

var books map[string]*Book

func init() {
	books = make(map[string]*Book, 100)
}

func GetBookByName(ctx context.Context, name string) interface{} {
	book, found := books[name]
	if found {
		return book
	}
	return nil
}

func GetBooks(ctx context.Context, limit int) interface{} {
	listLen := len(books)
	if limit > 0 && limit < len(books) {
		listLen = limit
	}

	list := make([]*Book, listLen)

	counter := 0
	for _, b := range books {
		if limit > 0 && counter > limit {
			break
		}
		list[counter] = b
		counter++
	}

	return list
}

// TODO do not return error
// TODO return book
func InsertBook(ctx context.Context, book Book) error {
	books[book.Name] = &book
	return nil
}

// TODO do not return error
// TODO return toUpd
func UpdateBook(ctx context.Context, updBook Book) error {
	toUpd, found := books[updBook.Name]
	if found {
		if updBook.Price > 0.0 {
			toUpd.Price = updBook.Price
		}

		if updBook.Description != "" {
			toUpd.Description = updBook.Description
		}
	} else {
		return InsertBook(ctx, updBook)
	}
	return nil
}

// TODO do not return error
func DeleteBook(ctx context.Context, name string) error {
	delete(books, name)
	return nil
}

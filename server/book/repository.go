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

func GetBookList(ctx context.Context, limit int) interface{} {
	list := make([]*Book, len(books))

	counter := 0
	for _, b := range books {
		if counter > limit {
			break
		}
		list[counter] = b
		counter++
	}

	return list
}

func InsertBook(ctx context.Context, book Book) error {
	books[book.Name] = &book
	return nil
}

func UpdateBook(ctx context.Context, book Book) error {
	toUpd, found := books[book.Name]
	if found {
		toUpd.Price = book.Price
		toUpd.Description = book.Description
	} else {
		return InsertBook(ctx, book)
	}
	return nil
}

func DeleteBook(ctx context.Context, name string) error {
	delete(books, name)
	return nil
}

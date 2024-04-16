package main

import "fmt"

// Book represents a book in a bookshelf.
type Book struct {
	Name string
}

// Iterator defines the interface for iterating over a collection.
type Iterator interface {
	HasNext() bool
	Next() *Book
}

// BookShelfIterator Aggregate defines the interface for a collection that can be iterated over.
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

// HasNext returns true if there are more elements to iterate over.
func (it *BookShelfIterator) HasNext() bool {
	return it.index < it.bookShelf.GetLength()
}

// Next returns the next book in the bookshelf.
func (it *BookShelfIterator) Next() *Book {
	if !it.HasNext() {
		return nil
	}
	book := it.bookShelf.GetBookAt(it.index)
	it.index++
	return book
}

// Aggregate defines the interface for a collection that can be iterated over.
type Aggregate interface {
	Iterator() Iterator
}

// BookShelf represents a collection of books.
type BookShelf struct {
	books []*Book
}

// Add appends a book to the bookshelf.
func (bs *BookShelf) Add(book *Book) {
	bs.books = append(bs.books, book)
}

// GetBookAt returns the book at the specified index.
func (bs *BookShelf) GetBookAt(index int) *Book {
	return bs.books[index]
}

// GetLength returns the number of books in the bookshelf.
func (bs *BookShelf) GetLength() int {
	return len(bs.books)
}

// Iterator returns a new iterator for the bookshelf.
func (bs *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{bookShelf: bs}
}

func main() {
	bookShelf := &BookShelf{}
	bookShelf.Add(&Book{Name: "Around the World in 80 Days"})
	bookShelf.Add(&Book{Name: "Bible"})
	bookShelf.Add(&Book{Name: "Cinderella"})
	bookShelf.Add(&Book{Name: "Daddy-Long-Legs"})

	it := bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Println(book.Name)
	}
}

package main

import "go_design_pattern/design_pattern/iterator"

func main() {
	testIterator()
}

func testIterator() {
	bookShelf := new(iterator.BookShelf)
	bookShelf.AppendBook(iterator.NewBook("Around the World in 80 Days"))
	bookShelf.AppendBook(iterator.NewBook("Bible"))
	bookShelf.AppendBook(iterator.NewBook("Cinderella"))
	bookShelf.AppendBook(iterator.NewBook("Daddy-Long-Legs"))

	it := bookShelf.Iterator()
	for it.HasNext() {
		book := it.Next().(*iterator.Book)
		println(book.GetName())
	}
}

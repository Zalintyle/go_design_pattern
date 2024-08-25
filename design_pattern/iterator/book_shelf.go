package iterator

type BookShelf struct {
	books []*Book
	last  int
}

func (bs *BookShelf) GetBookAt(index int) *Book {
	if bs != nil {
		return bs.books[index]
	}
	return nil
}

func (bs *BookShelf) AppendBook(book *Book) {
	if bs == nil {
		return
	}
	if bs.books == nil {
		bs.books = make([]*Book, 0)
	}
	bs.books = append(bs.books, book)
	bs.last++
}

func (bs *BookShelf) GetLength() int {
	if bs != nil {
		return bs.last
	}
	return 0
}

func (bs *BookShelf) Iterator() Iterator {
	if bs == nil {
		return nil
	}
	return NewBookShelfIterator(bs)
}

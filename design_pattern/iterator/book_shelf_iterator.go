package iterator

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	if bookShelf != nil {
		return &BookShelfIterator{
			bookShelf: bookShelf,
			index:     0,
		}
	}
	return nil
}

func (bsi *BookShelfIterator) HasNext() bool {
	if bsi != nil {
		if bsi.index < bsi.bookShelf.GetLength() {
			return true
		}
	}
	return false
}

func (bsi *BookShelfIterator) Next() interface{} {
	if bsi != nil {
		book := bsi.bookShelf.GetBookAt(bsi.index)
		bsi.index++
		return book
	}
	return nil
}

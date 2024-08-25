package iterator

type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{name: name}
}

func (b *Book) GetName() string {
	if b != nil {
		return b.name
	}
	return ""
}

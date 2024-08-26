package responsibility_chain

type Department interface {
	execute(*Patient)
	setNext(Department)
}

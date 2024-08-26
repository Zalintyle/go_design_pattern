package responsibility_chain

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		println("Medicine already given")
		m.next.execute(p)
		return
	}

	// handle the request
	println("Medical giving medicine")

	p.medicineDone = true
	if m.next == nil {
		println("All done")
		return
	}
	m.next.execute(p)
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

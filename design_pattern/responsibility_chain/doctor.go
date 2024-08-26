package responsibility_chain

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		println("Doctor checkup already done")
		d.next.execute(p)
		return
	}

	// handle the request
	println("Doctor checking patient")

	p.doctorCheckUpDone = true
	if d.next == nil {
		println("All done")
		return
	}
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

package responsibility_chain

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		println("Payment already done")
		c.next.execute(p)
		return
	}

	// handle the request
	println("Cashier collecting payment")

	p.paymentDone = true
	if c.next == nil {
		println("All done")
		return
	}
	c.next.execute(p)
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

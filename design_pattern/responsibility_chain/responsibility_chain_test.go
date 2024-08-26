package responsibility_chain

import "testing"

func TestResponsibilityChain(t *testing.T) {
	cashier := &Cashier{}
	cashier.setNext(nil)

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	reception.execute(patient)

	// Output:
	// Reception registering patient
	// Doctor checking patient
	// Medical giving medicine
	// Cashier collecting payment
	// All done
}

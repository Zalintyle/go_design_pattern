package demo

type Director struct {
	builder InterfaceBuilder
}

func NewDirector(b InterfaceBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b InterfaceBuilder) {
	d.builder = b
}

func (d *Director) Construct() Car {
	d.builder.SetSeatsType()
	d.builder.SetEngineType()
	d.builder.SetNumber()
	return d.builder.GetCar()
}

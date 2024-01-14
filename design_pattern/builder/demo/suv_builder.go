package demo

type SuvBuilder struct {
	Car
}

func (s *SuvBuilder) SetSeatsType() {
	s.SeatsType = "suv seats"
}

func (s *SuvBuilder) SetEngineType() {
	s.EngineType = "suv engine"
}

func (s *SuvBuilder) SetNumber() {
	s.Number = 6
}

func (s *SuvBuilder) GetCar() Car {
	return s.Car
}

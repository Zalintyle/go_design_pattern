package demo

type MpvBuilder struct {
	Car
}

func NewMpvBuilder() *MpvBuilder {
	return &MpvBuilder{}
}

func (m *MpvBuilder) SetSeatsType() {
	m.SeatsType = "mpv seats"
}

func (m *MpvBuilder) SetEngineType() {
	m.EngineType = "mpv engine"
}

func (m *MpvBuilder) SetNumber() {
	m.Number = 8
}

func (m *MpvBuilder) GetCar() Car {
	return m.Car
}

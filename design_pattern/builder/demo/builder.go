package demo

type InterfaceBuilder interface {
	SetSeatsType()
	SetEngineType()
	SetNumber()
	GetCar() Car
}

func GetBuilder(builderType string) InterfaceBuilder {
	switch builderType {
	case "mpv":
		return NewMpvBuilder()
	case "suv":
		return NewSuvBuilder()
	default:
		return nil
	}
}

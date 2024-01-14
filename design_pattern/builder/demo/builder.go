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
		return &MpvBuilder{}
	case "suv":
		return &SuvBuilder{}
	default:
		return nil
	}
}

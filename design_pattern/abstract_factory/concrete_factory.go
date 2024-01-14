package abstract_factory

type XiaomiFactory struct{}

func (x *XiaomiFactory) MakePhone() AbstractPhone {
	return &XiaomiPhone{
		Phone{
			color: "white",
			size:  1,
		},
	}
}

func (x *XiaomiFactory) MakeComputer() AbstractComputer {
	return &XiaomiComputer{
		Computer{
			color: "yellow",
			size:  3,
		},
	}
}

type LenovoFactory struct{}

func (x *LenovoFactory) MakePhone() AbstractPhone {
	return &LenovoPhone{
		Phone{
			color: "red",
			size:  2,
		},
	}
}

func (x *LenovoFactory) MakeComputer() AbstractComputer {
	return &LenovoComputer{
		Computer{
			color: "blue",
			size:  4,
		},
	}
}

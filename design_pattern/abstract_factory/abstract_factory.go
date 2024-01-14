package abstract_factory

import (
	"fmt"
)

// AbstractElectronicFactory 抽象工厂接口
type AbstractElectronicFactory interface {
	MakePhone() AbstractPhone       // 创建抽象商品1-Phone
	MakeComputer() AbstractComputer // 创建抽象商品2-Computer
}

func GetElectronicFactory(brand string) (AbstractElectronicFactory, error) {
	switch brand {
	case "Xiaomi":
		return &XiaomiFactory{}, nil
	case "Lenovo":
		return &LenovoFactory{}, nil
	default:
		return nil, fmt.Errorf("unknown brand:%v", brand)
	}
}

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

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

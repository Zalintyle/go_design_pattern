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

type XiaomiFactory struct {
}

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

type LenovoFactory struct {
}

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

// AbstractPhone 抽象产品接口-手机
type AbstractPhone interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

type Phone struct {
	color string
	size  int
}

func (p *Phone) SetColor(color string) {
	p.color = color
}

func (p *Phone) SetSize(size int) {
	p.size = size
}

func (p *Phone) GetColor() string {
	return p.color
}

func (p *Phone) GetSize() int {
	return p.size
}

// XiaomiPhone 间接实现了AbstractPhone接口，组合Phone
type XiaomiPhone struct {
	Phone
}

// LenovoPhone 间接实现了AbstractPhone接口
type LenovoPhone struct {
	Phone
}

// AbstractComputer 抽象产品接口-电脑
type AbstractComputer interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

type Computer struct {
	color string
	size  int
}

// XiaomiComputer 间接实现了AbstractComputer接口，组合Computer
type XiaomiComputer struct {
	Computer
}

// LenovoComputer 间接实现了AbstractComputer接口
type LenovoComputer struct {
	Computer
}

func (c *Computer) SetColor(color string) {
	c.color = color
}

func (c *Computer) SetSize(size int) {
	c.size = size
}

func (c *Computer) GetColor() string {
	return c.color
}

func (c *Computer) GetSize() int {
	return c.size
}

package factory_method_pattern

import "fmt"

type IClothes interface {
	setName(name string)
	setSize(size int)
	GetName() string
	GetSize() int
}

// clothes 实现了 IClothes 接口
type clothes struct {
	name string
	size int
}

func (c *clothes) setName(name string) {
	c.name = name
}

func (c *clothes) setSize(size int) {
	c.size = size
}

func (c *clothes) GetName() string {
	return c.name
}

func (c *clothes) GetSize() int {
	return c.size
}

// PEAK 嵌入了 clothes, 间接实现了 IClothes 接口
type PEAK struct {
	clothes
}

func newPEAK() IClothes {
	return &PEAK{
		clothes: clothes{
			name: "PEAK clothes",
			size: 1,
		},
	}
}

type ANTA struct {
	clothes
}

func newANTA() IClothes {
	return &ANTA{
		clothes: clothes{
			name: "ANTA clothes",
			size: 2,
		},
	}
}

// MakeClothes 工厂方法, 根据实参类型生产不同品牌的服装
func MakeClothes(clothesType string) (IClothes, error) {
	switch clothesType {
	case "PEAK":
		return newPEAK(), nil
	case "ANTA":
		return newANTA(), nil
	default:
		return nil, fmt.Errorf("unknown clothes type: %s", clothesType)
	}
}

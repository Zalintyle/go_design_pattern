package factory_method

import "fmt"

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

// ANTA 嵌入了 clothes, 间接实现了 IClothes 接口
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

// MakeClothes 工厂方法, 根据实参类型生产不同品牌的服装, 也可与具体工厂类绑定
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

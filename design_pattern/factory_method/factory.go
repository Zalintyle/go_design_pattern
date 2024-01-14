package factory_method

// IClothes 工厂接口
type IClothes interface {
	setName(name string)
	setSize(size int)
	GetName() string
	GetSize() int
}

package factory_method

// clothes 实现了 IClothes 接口, 为具体工厂类
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

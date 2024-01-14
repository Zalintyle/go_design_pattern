package abstract_factory

// Computer implement AbstractComputer
type Computer struct {
	color string
	size  int
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

// XiaomiComputer 间接实现了AbstractComputer接口，组合Computer
type XiaomiComputer struct {
	Computer
}

// LenovoComputer 间接实现了AbstractComputer接口
type LenovoComputer struct {
	Computer
}

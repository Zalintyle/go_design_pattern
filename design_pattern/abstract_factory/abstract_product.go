package abstract_factory

// AbstractPhone 抽象产品接口-手机
type AbstractPhone interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

// Phone implement AbstractPhone
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

// Computer implement AbstractComputer
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

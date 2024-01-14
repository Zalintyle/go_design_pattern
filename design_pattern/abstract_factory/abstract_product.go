package abstract_factory

// AbstractPhone 抽象产品接口-手机
type AbstractPhone interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

// AbstractComputer 抽象产品接口-电脑
type AbstractComputer interface {
	SetColor(color string)
	SetSize(size int)
	GetColor() string
	GetSize() int
}

package define

type ConcreteBuilder struct {
	result Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		result: Product{},
	}
}

func (cb *ConcreteBuilder) Build() {
	cb.result = Product{}
}

func (cb *ConcreteBuilder) GetResult() Product {
	return cb.result
}

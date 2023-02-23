package services

import "calculator/domain/valueobject"

type Adder struct{}

func NewAdder() *Adder {
	return new(Adder)
}

func (s *Adder) Add(augend, addend valueobject.Number) valueobject.Number {
	return augend.Plus(addend)
}

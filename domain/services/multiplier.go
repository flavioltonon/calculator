package services

import "calculator/domain/valueobject"

type Multiplier struct{}

func NewMultiplier() *Multiplier {
	return new(Multiplier)
}

func (s *Multiplier) Multiply(multiplier, multiplicand valueobject.Number) valueobject.Number {
	return multiplier.Times(multiplicand)
}

package services

import "calculator/domain/valueobject"

type Subtractor struct{}

func NewSubtractor() *Subtractor {
	return new(Subtractor)
}

func (s *Subtractor) Subtract(minuend, subtrahend valueobject.Number) valueobject.Number {
	return minuend.Minus(subtrahend)
}

package services

import (
	"calculator/domain/usecases"
	"calculator/domain/valueobject"
)

type Calculator struct {
	adder      usecases.Adder
	subtractor usecases.Subtractor
	multiplier usecases.Multiplier
	divider    usecases.Divider
}

func NewCalculator(
	adder usecases.Adder,
	subtractor usecases.Subtractor,
	multiplier usecases.Multiplier,
	divider usecases.Divider) *Calculator {
	return &Calculator{
		adder:      adder,
		subtractor: subtractor,
		multiplier: multiplier,
		divider:    divider,
	}
}

func (s *Calculator) Add(augend, addend valueobject.Number) valueobject.Number {
	return s.adder.Add(augend, addend)
}

func (s *Calculator) Subtract(minuend, subtrahend valueobject.Number) valueobject.Number {
	return s.subtractor.Subtract(minuend, subtrahend)
}

func (s *Calculator) Multiply(multiplier, multiplicand valueobject.Number) valueobject.Number {
	return s.multiplier.Multiply(multiplier, multiplicand)
}

func (s *Calculator) Divide(numerator, denominator valueobject.Number) valueobject.Number {
	return s.divider.Divide(numerator, denominator)
}

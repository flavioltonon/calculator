package services

import "calculator/domain/valueobject"

type Divider struct{}

func NewDivider() *Divider {
	return new(Divider)
}

func (s *Divider) Divide(numerator, denominator valueobject.Number) valueobject.Number {
	return numerator.DividedBy(denominator)
}

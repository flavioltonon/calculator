package usecases

import "calculator/domain/valueobject"

type Divider interface {
	Divide(numerator, denominator valueobject.Number) valueobject.Number
}

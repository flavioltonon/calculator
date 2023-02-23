package usecases

import "calculator/domain/valueobject"

type Multiplier interface {
	Multiply(multiplier, multiplicand valueobject.Number) valueobject.Number
}

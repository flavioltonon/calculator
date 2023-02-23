package usecases

import "calculator/domain/valueobject"

type Subtractor interface {
	Subtract(minuend, subtrahend valueobject.Number) valueobject.Number
}

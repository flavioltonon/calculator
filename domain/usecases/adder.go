package usecases

import "calculator/domain/valueobject"

type Adder interface {
	Add(augend, addend valueobject.Number) valueobject.Number
}

package dto

import "calculator/domain/valueobject"

type Calculation struct {
	Result float64 `json:"result"`
}

func NewCalculation(result valueobject.Number) Calculation {
	return Calculation{
		Result: result.Value(),
	}
}

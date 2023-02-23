package valueobject

import "math"

type Number float64

func (n Number) Value() float64 {
	return float64(n)
}

func (n Number) epsilon() float64 {
	return 1e-6
}

func (n Number) Plus(ref Number) Number {
	return n + ref
}

func (n Number) Minus(ref Number) Number {
	return n - ref
}

func (n Number) Times(ref Number) Number {
	return n * ref
}

func (n Number) DividedBy(ref Number) Number {
	return n / ref
}

func (n Number) Equal(ref Number) bool {
	delta := n.Minus(ref)
	return math.Abs(delta.Value()) <= n.epsilon()
}

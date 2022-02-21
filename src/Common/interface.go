package common

type Coster interface {
	Cost() float64
}

func DisplayCost(c Coster) float64 {
	return c.Cost()
}

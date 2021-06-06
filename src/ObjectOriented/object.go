package object

import (
	"fmt"
)

type quantity int
type costCalculator func(quantity, float64) float64

func describe(q quantity, price float64, c costCalculator) {
	fmt.Printf("quantity: %d, price: %0.0f, cost: %0.0f\n", q, price, c(q, price))
}

func signatureCustomType() {
	var offBy10Percent, offBy1000Won costCalculator

	offBy10Percent = func(q quantity, price float64) float64 {
		return float64(q) * price * 0.9
	}
	offBy1000Won = func(q quantity, price float64) float64 {
		return float64(q)*price - 1000
	}

	describe(3, 10000, offBy10Percent)
	describe(3, 10000, offBy1000Won)
}

func (q quantity) greaterThan(i int) bool { return int(q) > i }

func (q *quantity) increment() { *q++ }

func (q *quantity) decrement() { *q-- }

func referencedReceiver() {
	q := quantity(3)
	q.increment()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))
	q.decrement()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))
}

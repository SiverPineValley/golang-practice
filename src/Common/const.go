package common

import (
	"fmt"
	"reflect"
	"strings"
)

/*
	도형 관련 const
*/
type Rect struct{ Width, Height float64 }

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r *Rect) Resize(w, h float64) {
	r.Width += w
	r.Height += h
}

func (r *Rect) Show() {
	fmt.Printf("width: %f, height: %f\n", r.Width, r.Height)
}

type Circle struct{ Radius float64 }

func (c *Circle) Show() {
	fmt.Printf("radius: %f\n", c.Radius)
}

func Display(s interface{ Show() }) {
	s.Show()
}

/*
	가격 관련 const
*/

type Item struct {
	Name     string  "json:`name` 상품 이름"
	Price    float64 "json:`price` 상품 가격"
	Quantity int     "json:`quantity` 상품 수량"
}

func (i Item) Cost() float64 {
	return i.Price * float64(i.Quantity)
}

func (i Item) String() string {
	return fmt.Sprintf("[%s] %.0f", i.Name, i.Cost())
}

type DiscountItem struct {
	Item
	DiscountRate float64
}

func (d DiscountItem) Cost() float64 {
	return d.Item.Cost() * (1.0 - d.DiscountRate/100)
}

func (d DiscountItem) String() string {
	return fmt.Sprintf("%s => %.0f(%.0f%s DC)", d.Name, d.Cost(), d.DiscountRate, "%")
}

type RentalPeriod int

const (
	Days RentalPeriod = iota
	Weeks
	Months
)

func (p RentalPeriod) Todays() int {
	switch p {
	case Weeks:
		return 7
	case Months:
		return 30
	default:
		return 1
	}
}
type Rental struct {
	Name         string
	FeePerDay    float64
	PeriodLength int
	RentalPeriod
}

func (r Rental) Cost() float64 {
	return r.FeePerDay * float64(r.Todays() * r.PeriodLength)
}

func (r Rental) String() string {
	return fmt.Sprintf("[%s] %.0f", r.Name, r.Cost())
}

type Items []Coster

func (it Items) Cost() (c float64) {
	for _, t := range it {
		c += t.Cost()
	}
	return
}

func (it Items) String() string {
	var s []string
	for _, t := range it {
		s = append(s, fmt.Sprint(t))
	}
	return fmt.Sprintf("%d items. total: %.0f\n\t- %s", len(it), it.Cost(), strings.Join(s, "\n\t- "))
}

func CheckCosterTypeAssertion(v interface{}) (ok bool) {
	_, ok = v.(Coster)
	return
}

func CheckType(v interface{}) {
	if v == nil {
		fmt.Printf("%v is nil\n", v)
		return
	}

	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Bool:
		fmt.Printf("%t is a bool\n", v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("%d is an int\n", v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Printf("%d is an unsigned int\n", v)
	case reflect.Float32, reflect.Float64:
		fmt.Printf("%f is float\n", v)
	case reflect.Complex64, reflect.Complex128:
		fmt.Printf("%f is a complex\n", v)
	case reflect.String:
		fmt.Printf("%s is a string\n", v)
	case reflect.Struct:
		fmt.Printf("%v is a struct\n", v)
	case reflect.Func:
		fmt.Printf("%v is a function\n", v)
	case reflect.Array:
		fmt.Printf("%v is a array\n", v)
	case reflect.Slice:
		fmt.Printf("%v is a slice\n", v)
	case reflect.Map:
		fmt.Printf("%v is a map\n", v)
	default:
		fmt.Printf("%v is unknown type\n", v)
	}
}
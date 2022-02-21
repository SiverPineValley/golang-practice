package object

import (
	"fmt"
	common "golang-practice/src/Common"
	"reflect"
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

func methodToFunction() {
	r := common.Rect{3, 4}
	fmt.Println("area :", r.Area()) // area: 12
	r.Resize(10, 10)
	fmt.Println("area :", r.Area()) // area: 182

	// Area() 메서드의 함수 표혀식
	areaFunc := common.Rect.Area
	resizeFunc := (*common.Rect).Resize

	fmt.Println("area: ", areaFunc(r))
	resizeFunc(&r, -10, -10)
	fmt.Println("area: ", areaFunc(r))
}

func structTag() {
	tType := reflect.TypeOf(common.Item{})
	for i := 0; i < tType.NumField(); i++ {
		fmt.Println(tType.Field(i).Tag)
	}
}

func structMethodReusing() {
	shoes := common.Item{"Women's Walking Shoes", 30000, 2}
	eventShoes := common.DiscountItem{
		common.Item{"Sports Shoes", 50000, 3},
		10.00,
	}

	fmt.Println(shoes.Cost())           // 60000
	fmt.Println(eventShoes.Cost())      // 135000
	fmt.Println(eventShoes.Item.Cost()) // 150000
}

func anonymousInterface() {
	r := common.Rect{3, 4}
	c := common.Circle{2.5}

	common.Display(&r)
	common.Display(&c)
}

func interfacePolymorphism() {
	item := common.Item{"Sports Shoes", 30000, 2}
	eventItem := common.DiscountItem{common.Item{"Sports Shoes (DC)", 30000, 3}, 10.00}

	fmt.Println(common.DisplayCost(item))
	fmt.Println(common.DisplayCost(eventItem))

	shirts := common.Item{"Men's Slim-Fit Shirt",25000,3}
	video := common.Rental{"Interstellar",1000,2, common.Weeks}

	fmt.Println(common.DisplayCost(shirts))
	fmt.Println(common.DisplayCost(video))
}

func genericCollection() {
	shirt := common.Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := common.Rental{"Interstellar", 1000, 3, common.Days}
	eventShoes := common.DiscountItem{common.Item{"Women's Walking Shoes", 50000, 3}, 10.00}
	items := common.Items{shirt, video, eventShoes}
	fmt.Println(common.DisplayCost(items))
}

func fmtStringer() {
	shirt := common.Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := common.Rental{"Interstellar", 1000, 3, common.Days}
	eventShoes := common.DiscountItem{common.Item{"Women's Walking Shoes", 50000, 3}, 10.00}
	items := common.Items{shirt, video, eventShoes}

	fmt.Println(shirt)
	fmt.Println(video)
	fmt.Println(eventShoes)
	fmt.Println(items)
}

func typeAssertion() {
	shirt := common.Item{"Men's Slim-Fit Shirt", 25000, 3}

	fmt.Println("Type Assertion Test from shirt and int to Coaster")
	fmt.Println(shirt.String(),": " , common.CheckCosterTypeAssertion(shirt))
	fmt.Println("100: ", common.CheckCosterTypeAssertion(100))
}

func typeAssertionBySwitch() {
	shirt := common.Item{"Men's Slim-Fit Shirt", 25000, 3}

	common.CheckType(3)
	common.CheckType(1.546)
	common.CheckType(complex(1, 5))
	common.CheckType(true)
	common.CheckType("s")
	common.CheckType(shirt)
	common.CheckType([]int{0})
	common.CheckType([]int{0,1,2}[0:1])
	common.CheckType(map[string]string{})
	common.CheckType(struct{}{})
	common.CheckType(nil)
	common.CheckType(typeAssertion)
}
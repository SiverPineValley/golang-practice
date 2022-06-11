package Package

import (
	"container/list"
	"fmt"
	common "golang-practice/src/Common"
	"reflect"
)

func OsShow() {
	showOS()
}

func ReflectionValueCheck() {
	u := common.User{"Jang", 1}
		uType := reflect.TypeOf(u)

		if fName, ok := uType.FieldByName("Name"); ok {
				fmt.Println(fName.Type, fName.Name, fName.Tag)
			}

		if fId, ok := uType.FieldByName("Id"); ok {
				fmt.Println(fId.Type, fId.Name, fId.Tag)
			}
	}

func ReflectionValueChange() {
	x := 1
	if v := reflect.ValueOf(x); v.CanSet() {
		v.SetInt(2) // 호출되지 않음
	}

	fmt.Println(x) // 1

	v := reflect.ValueOf(&x)
	p := v.Elem()
	p.SetInt(3)
	fmt.Println(x) // 3
}

func ReflectionDynamicFunctionCall() {
	caption := "go is an open source programming language"

	// TitleCase를 바로 호출
	title := common.TitleCase(caption)
	fmt.Println(title)

	// TitleCase를 동적으로 호출
	titleFuncValue := reflect.ValueOf(common.TitleCase)
	values := titleFuncValue.Call([]reflect.Value{reflect.ValueOf(caption)})
	title = values[0].String()
	fmt.Println(title)
}

func ReflectionLen() {
	a := list.New()
	b := list.New()
	b.PushFront(0.5)
	c := map[string]int{"A": 1, "B": 2}
	d := "one"
	e := []int{5, 0, 4, 1}

	fmt.Println(common.Len(a), common.Len(b), common.Len(c), common.Len(d), common.Len(e))
}
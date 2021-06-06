package main

import (
	"fmt"
	object "golang-practice/src/ObjectOriented"
	types "golang-practice/src/Types"
	bg "golang-practice/src/basicGrammar"
	_ "os"
)

func main() {
	var chapter string

	fmt.Printf("1. 기본 문법\n")
	fmt.Printf("2. 데이터 타입\n")
	fmt.Printf("3. 객체 지향 프로그래밍\n")
	fmt.Printf("입력하세요: ")
	fmt.Scanln(&chapter)

	switch chapter {
	case "1":
		bg.BasicGrammar()
	case "2":
		types.Types()
	case "3":
		object.Object()
	}

}

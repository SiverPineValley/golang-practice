package main

import (
	"fmt"
	bg "golang-practice/src/basicGrammar"
	_ "os"
)

func main() {
	var chapter string

	fmt.Printf("1. 기본 문법\n")
	fmt.Printf("입력하세요: ")
	fmt.Scanln(&chapter)

	switch chapter {
	case "1":
		bg.BasicGrammar()
	}

}

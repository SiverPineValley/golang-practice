package basicGrammar

import (
	"fmt"
	"strings"
)

// 가변인자 예시 함수
func multiParams(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Printf(strings[i] + " ")
	}
}

// Call by value
func callbyValue(i int) {
	i = i + 1
}

// Call by reference
func callbyReference(i *int) {
	*i = *i + 1
}

// Defer Example
func deferFunction() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d번째 호출된 Defer\n", i+1)
	}
}

func enter(s string) string {
	fmt.Println("entering: ", s)
	return s
}

func leave(s string) {
	fmt.Println("leaving: ", s)
}

func a() {
	defer leave(enter("a"))
	fmt.Println("in a")
}

func b() {
	defer leave(enter("b"))
	fmt.Println("in b")
	a()
}

// Closure Example
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func closureConst(value int) string {
	mappingTable := map[int]string{
		0: "SKT",
		1: "LG U+",
		2: "KT",
		3: "알뜰 SKT",
		4: "알뜰 LG U+",
		5: "알뜰 KT",
		6: "없음",
	}

	if value > 6 {
		value = 6
	}

	return mappingTable[value]
}

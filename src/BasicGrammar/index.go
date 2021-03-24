package basicGrammar

import (
	"fmt"
	"strings"
	"unicode"
)

func BasicGrammar() {
	fmt.Println("1. 레이블 예시")
	useLable1()
	useLable2()

	fmt.Println("\n2. 가변 매개변수 예시")
	multiParams("Variable", "Params", "example")

	fmt.Println("\n3. Call by Value, Reference")
	var i int = 1
	callbyValue(i)
	fmt.Println("Call by Value: ", i)
	callbyReference(&i)
	fmt.Println("Call by Rererence: ", i)

	fmt.Println("\n4. defer")
	deferFunction()
	b()

	fmt.Println("\n5. Closure")
	addZip := makeSuffix(".zip")
	addTgz := makeSuffix(".tar.gz")
	fmt.Println(addTgz("go1.5.1.src"))
	fmt.Println(addZip("go1.5.1.windows-amd64"))
	fmt.Println(closureConst(0))
	fmt.Println(closureConst(4))
	fmt.Println(closureConst(12))

	f := func(c rune) bool {
		return unicode.Is(unicode.Hangul, c)
	}

	fmt.Println(strings.IndexFunc("\nHello, 월드", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))

}

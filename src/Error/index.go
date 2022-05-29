package Error

import "fmt"

func ErrorProcessing() {
	fmt.Println("1. recover")
	Recover()

	fmt.Println("\n2. protect")
	Protect()

	fmt.Println("\n3. 클로저로 에러 처리")
	CloserError()
}

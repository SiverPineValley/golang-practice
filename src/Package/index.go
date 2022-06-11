package Package

import "fmt"

func Package() {
	fmt.Println("\n1. 운영체제에 종속적인 코드 처리")
	OsShow()

	fmt.Println("\n2. Reflection 값 확인")
	ReflectionValueCheck()

	fmt.Println("\n3. Reflection 값 변경")
	ReflectionValueChange()

	fmt.Println("\n4. Reflection 함수/메서드 동적 호출")
	ReflectionDynamicFunctionCall()

	fmt.Println("\n5. Reflection 함수/메서드 동적 호출 - 2")
	ReflectionLen()

	fmt.Println("\n6. 테스트")
	ReflectionLen()
}

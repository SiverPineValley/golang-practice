package object

import (
	"fmt"
)

func Object() {
	fmt.Println("1. 함수 서명을 사용자 정의 타입으로 사용")
	signatureCustomType()

	fmt.Println("\n2. 참조에 의한 리시버")
	referencedReceiver()

	fmt.Println("\n3. 메서드의 함수 표현식")
	methodToFunction()

	fmt.Println("\n4. 구조체의 태그")
	structTag()

	fmt.Println("\n5. 구조체 메서드 재사용")
	structMethodReusing()

	fmt.Println("\n6. 익명 인터페이스")
	anonymousInterface()

	fmt.Println("\n7. 인터페이스를 사용한 다형성 구현")
	interfacePolymorphism()

	fmt.Println("\n8. Generic Collection")
	genericCollection()

	fmt.Println("\n9. fmt.Stringer 활용")
	fmtStringer()

	fmt.Println("\n10. Type Assertion")
	typeAssertion()

	fmt.Println("\n11. Type Assertion By Switch")
	typeAssertionBySwitch()
}

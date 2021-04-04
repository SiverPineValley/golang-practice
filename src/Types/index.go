package types

import (
	"fmt"
)

func Types() {
	fmt.Println("1. 진수")
	NumbericSystems()

	fmt.Println("\n2. 문자 표기")
	ByteAndRune()

	fmt.Println("\n3. 복소수")
	ComplexType()

	fmt.Println("\n4. 이스케이프")
	Escape()

	fmt.Println("\n5. 문자열 길이")
	Len()

	fmt.Println("\n6. 문자열 확인")
	Check()

	fmt.Println("\n7. 문자열 Concat")
	Concat()

	fmt.Println("\n8. Slice")
	Slice()

	fmt.Println("\n9. 정렬")
	Sort()
}

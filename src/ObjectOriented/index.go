package object

import (
	"fmt"
)

func Object() {
	fmt.Println("1. 함수 서명을 사용자 정의 타입으로 사용")
	signatureCustomType()

	fmt.Println("\n2. 참조에 의한 리시버")
	referencedReceiver()
}

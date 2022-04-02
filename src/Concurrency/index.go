package Concurrency

import "fmt"

func Concurrency() {
	fmt.Println("1. 고루틴")
	goRoutine()

	fmt.Println("\n\n2. 채널")
	channel()

	fmt.Println("\n\n3. 버퍼드 채널")
	bufferedChannel()

	fmt.Println("\n\n4. select")
	selectChannel()

	fmt.Println("\n\n5. 저수준 제어")
	fmt.Println("\n5.1. mutex 미사용")
	notUsingmutex()
	fmt.Println("\n5.2. mutex 사용")
	mutex()
	fmt.Println("\n5.3. sync.Once 사용")
	syncOnce()
	fmt.Println("\n5.4. sync.WaitGroup 사용")
	waitGroup()
	fmt.Println("\n5.5. 원자성을 보장하는 연산")
	atomic()
}

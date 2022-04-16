package Concurrency

import (
	"fmt"
	common "golang-practice/src/Common"
	"runtime"
	"sync"
	"time"
)

func goRoutine() {
	fmt.Println("main 함수 시작: ", time.Now())

	go common.Long()
	go common.Short()

	time.Sleep(5 * time.Second)
	fmt.Println("main 함수 종료: ", time.Now())
}

func channel() {
	fmt.Println("main 함수 시작: ", time.Now())

	var done chan bool
	done = make(chan bool)
	go common.LongChan(done)
	go common.ShortChan(done)
	<-done
	<-done

	fmt.Println("main 함수 종료: ", time.Now())
}

func bufferedChannel() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	//c <- 3 -> 버퍼가 꽉 찼는데 메시지를 계속 전송해서 에러 발생
	go func() { c <- 3 }() // -> 고루틴은 c 채널에 메시지를 전송할 수 있을때까지 기다렸다가 채널에 첫 번째 값을 수신하는 즉시 채널에 값을 전송

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

func selectChannel() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func notUsingmutex() {
	// 모든 CPU 를 사용하게 함
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := common.Counter{I: 0}   // 카운터 생성
	done := make(chan struct{}) // 완료 신호 수신용 채널

	for i := 0; i < 1000; i++ {
		go func() {
			c.Increment()
			done <- struct{}{}
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	c.Display()
}

func mutex() {
	// 모든 CPU 를 사용하게 함
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := common.Counter{I: 0}   // 카운터 생성
	done := make(chan struct{}) // 완료 신호 수신용 채널

	for i := 0; i < 1000; i++ {
		go func() {
			c.IncrementMutex()
			done <- struct{}{}
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	c.Display()
}

func syncOnce() {
	// 모든 CPU 를 사용하게 함
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := common.Counter{I: 0}   // 카운터 생성
	done := make(chan struct{}) // 완료 신호 수신용 채널

	for i := 0; i < 1000; i++ {
		go func() {
			c.IncrementSyncOnce()
			done <- struct{}{}
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	c.Display()
}

func waitGroup() {
	// 모든 CPU 를 사용하게 함
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := common.Counter{I: 0} // 카운터 생성
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.IncrementMutex()
		}()
	}

	wg.Wait()

	c.Display()
}

func atomic() {
	// 모든 CPU 를 사용하게 함
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := common.Counter{I: 0} // 카운터 생성
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.IncrementAtomic()
		}()
	}

	wg.Wait()

	c.Display()
}

func timeout() {
	quit := make(chan struct{})
	done := common.Process(quit)
	timeout := time.After(1 * time.Second)

	select {
	case d := <-done:
		fmt.Println(d)
	case <-timeout:
		quit <- struct{}{}
		fmt.Println("Time out!!")
	}

	return
}

func sharedMap() {
	m := common.NewMap()

	// Set item
	ok := m.Set("foo", "bar")
	fmt.Println(ok)

	// Get item
	t, ok := m.Get("foo")

	// Check if item exists
	if ok {
		bar := t.(string)
		fmt.Println("bar: ", bar)
	}

	// Count item
	fmt.Println("Count: ", m.Count())

	// Remove item
	ok = m.Remove("foo")
	fmt.Println(ok)

	// Count item
	fmt.Println("Count: ", m.Count())

	return
}

func pipeline() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	path := "/Users/parkjongin/workspace/go/src/golang-practice"
	pattern := ".go$"
	<-common.Show(common.Grep(pattern, common.Find(path)))
	return
}

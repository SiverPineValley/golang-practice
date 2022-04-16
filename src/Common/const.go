package common

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

/*
	도형 관련 const
*/
type Rect struct{ Width, Height float64 }

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r *Rect) Resize(w, h float64) {
	r.Width += w
	r.Height += h
}

func (r *Rect) Show() {
	fmt.Printf("width: %f, height: %f\n", r.Width, r.Height)
}

type Circle struct{ Radius float64 }

func (c *Circle) Show() {
	fmt.Printf("radius: %f\n", c.Radius)
}

func Display(s interface{ Show() }) {
	s.Show()
}

/*
	가격 관련 const
*/

type Item struct {
	Name     string  "json:`name` 상품 이름"
	Price    float64 "json:`price` 상품 가격"
	Quantity int     "json:`quantity` 상품 수량"
}

func (i Item) Cost() float64 {
	return i.Price * float64(i.Quantity)
}

func (i Item) String() string {
	return fmt.Sprintf("[%s] %.0f", i.Name, i.Cost())
}

type DiscountItem struct {
	Item
	DiscountRate float64
}

func (d DiscountItem) Cost() float64 {
	return d.Item.Cost() * (1.0 - d.DiscountRate/100)
}

func (d DiscountItem) String() string {
	return fmt.Sprintf("%s => %.0f(%.0f%s DC)", d.Name, d.Cost(), d.DiscountRate, "%")
}

type RentalPeriod int

const (
	Days RentalPeriod = iota
	Weeks
	Months
)

func (p RentalPeriod) Todays() int {
	switch p {
	case Weeks:
		return 7
	case Months:
		return 30
	default:
		return 1
	}
}

type Rental struct {
	Name         string
	FeePerDay    float64
	PeriodLength int
	RentalPeriod
}

func (r Rental) Cost() float64 {
	return r.FeePerDay * float64(r.Todays()*r.PeriodLength)
}

func (r Rental) String() string {
	return fmt.Sprintf("[%s] %.0f", r.Name, r.Cost())
}

type Items []Coster

func (it Items) Cost() (c float64) {
	for _, t := range it {
		c += t.Cost()
	}
	return
}

func (it Items) String() string {
	var s []string
	for _, t := range it {
		s = append(s, fmt.Sprint(t))
	}
	return fmt.Sprintf("%d items. total: %.0f\n\t- %s", len(it), it.Cost(), strings.Join(s, "\n\t- "))
}

func CheckCosterTypeAssertion(v interface{}) (ok bool) {
	_, ok = v.(Coster)
	return
}

func CheckType(v interface{}) {
	if v == nil {
		fmt.Printf("%v is nil\n", v)
		return
	}

	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Bool:
		fmt.Printf("%t is a bool\n", v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("%d is an int\n", v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Printf("%d is an unsigned int\n", v)
	case reflect.Float32, reflect.Float64:
		fmt.Printf("%f is float\n", v)
	case reflect.Complex64, reflect.Complex128:
		fmt.Printf("%f is a complex\n", v)
	case reflect.String:
		fmt.Printf("%s is a string\n", v)
	case reflect.Struct:
		fmt.Printf("%v is a struct\n", v)
	case reflect.Func:
		fmt.Printf("%v is a function\n", v)
	case reflect.Array:
		fmt.Printf("%v is a array\n", v)
	case reflect.Slice:
		fmt.Printf("%v is a slice\n", v)
	case reflect.Map:
		fmt.Printf("%v is a map\n", v)
	default:
		fmt.Printf("%v is unknown type\n", v)
	}
}

/*
	병행처리 관련 const
*/

func Long() {
	fmt.Println("long 함수 시작", time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println("long 함수 종료", time.Now())
}

func Short() {
	fmt.Println("short 함수 시작", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("short 함수 종료", time.Now())
}

func LongChan(done chan bool) {
	fmt.Println("long 함수 시작", time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println("long 함수 종료", time.Now())
	done <- true
}

func ShortChan(done chan bool) {
	fmt.Println("short 함수 시작", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("short 함수 종료", time.Now())
	done <- true
}

/*
	저수준 제어 관련 const
*/

type Counter struct {
	I    int64
	Mu   sync.Mutex
	once sync.Once
}

func (c *Counter) Increment() {
	c.I += 1
}

func (c *Counter) IncrementMutex() {
	c.Mu.Lock()
	c.I += 1
	c.Mu.Unlock()
}

func (c *Counter) IncrementSyncOnce() {
	c.once.Do(func() {
		c.I = -500
	})
	c.Mu.Lock()
	c.I += 1
	c.Mu.Unlock()
}

func (c *Counter) IncrementAtomic() {
	atomic.AddInt64(&c.I, 1)
}

func (c *Counter) Display() {
	fmt.Println(c.I)
}

/*
	병행 처리 활용 관련 함수
*/

const (
	set = iota
	get
	remove
	count
)

type SharedMap struct {
	m map[string]interface{}
	c chan Command
}

type Command struct {
	key    string
	value  interface{}
	action int
	result chan<- interface{}
}

func Process(quit <-chan struct{}) chan string {
	done := make(chan string)
	go func() {
		go func() {
			time.Sleep(10 * time.Second) // heavy job

			done <- "Complete !!"
		}()
		<-quit
		return
	}()

	return done
}

func (sm SharedMap) Set(k string, v interface{}) (r bool) {
	callback := make(chan interface{})
	sm.c <- Command{action: set, key: k, value: v, result: callback}
	r = (<-callback).(bool)
	return
}

func (sm SharedMap) Get(k string) (v interface{}, r bool) {
	callback := make(chan interface{})
	sm.c <- Command{action: get, key: k, result: callback}
	result := (<-callback).([2]interface{})
	v = result[0]
	r = result[1].(bool)
	return
}

func (sm SharedMap) Remove(k string) (r bool) {
	callback := make(chan interface{})
	sm.c <- Command{action: remove, key: k, result: callback}
	r = (<-callback).(bool)
	return
}

func (sm SharedMap) Count() (r int) {
	callback := make(chan interface{})
	sm.c <- Command{action: count, result: callback}
	r = (<-callback).(int)
	return
}

func (sm SharedMap) run() {
	for cmd := range sm.c {
		switch cmd.action {
		case set:
			sm.m[cmd.key] = cmd.value
			_, ok := sm.m[cmd.key]
			cmd.result <- ok
		case get:
			v, ok := sm.m[cmd.key]
			cmd.result <- [2]interface{}{v, ok}
		case remove:
			_, ok := sm.m[cmd.key]
			if !ok {
				cmd.result <- false
			} else {
				delete(sm.m, cmd.key)
				_, ok = sm.m[cmd.key]
				cmd.result <- !ok
			}
		case count:
			cmd.result <- len(sm.m)
		}
	}
}

func NewMap() SharedMap {
	sm := SharedMap{
		m: make(map[string]interface{}),
		c: make(chan Command),
	}
	go sm.run()
	return sm
}

const BUF_SIZE = 1000

var (
	workers = runtime.NumCPU()
)

func Find(path string) <-chan string {
	out := make(chan string, BUF_SIZE)
	done := make(chan struct{}, workers)
	for i := 0; i < workers; i++ {
		go func() {
			filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
				out <- file
				return nil
			})
			done <- struct{}{}
		}()
	}
	go func() {
		for i := 0; i < cap(done); i++ {
			<-done
		}
		close(out)
	}()

	return out
}

func Grep(pattern string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		regex, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println(err)
			return
		}

		for file := range in {
			if regex.MatchString(file) {
				out <- file
			}
		}
		close(out)
	}()
	return out
}

func Show(in <-chan string) <-chan struct{} {
	quit := make(chan struct{})
	go func() {
		for file := range in {
			c, err := lineCount(file)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("8d %s\n", c, file)
		}
		quit <- struct{}{}
	}()
	return quit
}

func lineCount(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return 0, err
	}

	if info.Mode().IsDir() {
		return 0, fmt.Errorf("%s is a directory", file)
	}

	count := 0
	buf := make([]byte, 1024*8)
	newLine := []byte{'\n'}

	for {
		c, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return 0, err
		}

		count += bytes.Count(buf[:c], newLine)

		if err == io.EOF {
			break
		}
	}
	return count, nil
}
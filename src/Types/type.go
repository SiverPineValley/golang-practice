package types

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// 숫자
func NumbericSystems() {
	a := 365   // 10진수
	b := 0555  // 8진수
	c := 0x16D // 16진수

	fmt.Println("10진수 365: ", a)
	fmt.Println("8진수 0555: ", b)
	fmt.Println("16진수 0x16D: ", c)
}

func ByteAndRune() {
	var bt1 byte = 65
	var bt2 byte = 0101
	var bt3 byte = 0x41

	var ru1 rune = 44032
	var ru2 rune = 0126000
	var ru3 rune = 0xAC00

	fmt.Printf("%c %c %c\n", bt1, bt2, bt3)
	fmt.Printf("%c %c %c\n", ru1, ru2, ru3)

	var ch1 byte = 'A'
	var ch2 rune = '가'

	fmt.Printf("%c %c\n", ch1, ch2)
}

func ComplexType() {
	c1 := 1 + 2i
	c2 := complex64(3 + 4i)
	c3 := complex(5, 6)

	fmt.Println(c1, real(c1), imag(c1))
	fmt.Println(c2, real(c2), imag(c2))
	fmt.Println(c3, real(c3), imag(c3))
}

// 문자열
func Escape() {
	path1 := "c:\\workspace\\go\\src"
	path2 := `c:\workspace\go\src`
	fmt.Println("Included Escape: ", path1)
	fmt.Println("Not Included Escape: ", path2)
}

func Len() {
	first := "Abcd123!23"
	second := "고랭배웁니다"
	rFirst := []rune(first)
	rSecond := []rune(second)

	fmt.Println("Length of "+first+": ", len(first))
	fmt.Println("Rune Length of "+first+": ", len([]rune(first)))
	fmt.Println("Length of "+second+": ", len(second))
	fmt.Println("Rune Length of "+second+": ", len([]rune(second)))

	fmt.Printf("s1: %c %c %c %c %c\n", rFirst[0], rFirst[1], rFirst[2], rFirst[3], rFirst[4])
	fmt.Printf("s2: %c %c %c %c %c\n", rSecond[0], rSecond[1], rSecond[2], rSecond[3], rSecond[4])
}

func Check() {
	five := []rune("5")
	B := []rune("b")
	Space := []rune(" ")

	fmt.Println("Is Number (5): ", unicode.IsNumber(five[0]))
	fmt.Println("Is Capital (b): ", unicode.IsUpper(B[0]))
	fmt.Println("Is Capital ( ): ", unicode.IsSpace(Space[0]))
}

func Concat() {
	fmt.Println(strings.Join([]string{"안녕하세요", ", 월드입니다"}, ""))
}

func insert(s []int, new []int, index int) []int {
	return append(s[:index], append(new, s[index:]...)...)
}

func Slice() {
	ns := [][]int{
		{1, 2, 3},
		{6, 7, 8},
		{8, 9, 10, 11},
	}

	result := make([]int, 0)

	for idx, value := range ns {
		switch idx {
		case 0:
			result = append(result, value...)
		case 1:
			result = append(result, 4, 5)
			result = append(result, value...)
		case 2:
			result = append(result, value[1:]...)

		}
		fmt.Printf("%d. len: %d, cap: %d, %v\n", idx, len(result), cap(result), result)
	}

	result = insert(result, []int{17, 18, 19}, 5)
	fmt.Printf("%d. len: %d, cap: %d, %v\n", 3, len(result), cap(result), result)
}

func Sort() {
	sli := []int{9, 8, 6, 5, 3}

	fmt.Printf("%v\n", sli)
	sort.Ints(sli)
	fmt.Printf("%v\n", sli)
}

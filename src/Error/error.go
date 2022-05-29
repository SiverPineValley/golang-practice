package Error

import (
	"fmt"
	common "golang-practice/src/Common"
)

func Recover() {
	fmt.Println("result: ", common.DivideRecover(1, 0))
}

func Protect() {
	common.Protect(func() {
		fmt.Println(common.Divide(1, 0))
	})
}

func CloserError() {
	fmt.Println(common.ErrorHandler(common.Divide)(4, 2))
	fmt.Println(common.ErrorHandler(common.Divide)(3, 0))
}

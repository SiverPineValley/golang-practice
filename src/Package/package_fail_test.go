package Package

import (
	common "golang-practice/src/Common"
	"testing"
)

func TestLenForChan(t *testing.T) {
	v := make(chan int)
	actual, expected := common.Len(v), 1
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}
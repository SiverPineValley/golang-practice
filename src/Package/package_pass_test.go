package Package

import (
	common "golang-practice/src/Common"
	"testing"
)

func TestLenForMap(t *testing.T) {
	v := map[string]int{"A": 1, "B": 2}
	actual, expected := common.Len(v), 2
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

func TestLenForString(t *testing.T) {
	v := "one"
	actual, expected := common.Len(v), 3
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

func TestLenForSlice(t *testing.T) {
	v := []int{5, 0, 4, 1}
	actual, expected := common.Len(v), 4
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

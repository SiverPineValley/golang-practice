package basicGrammar

import "fmt"

func useLable1() {
	x := 7
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 5}}

LOOP:
	for row, rowValue := range table {
		for col, colValue := range rowValue {
			if colValue == x {
				fmt.Printf("found %d(row:%d, col:%d)\n", x, row, col)
				break LOOP
			}
		}
	}
}

func useLable2() {
	x := 5
	table := [][]int{{1, 5, 9}, {2, 6, 5, 13}, {5, 3, 7, 4}}

LOOPCONTINUE:
	for row, rowValue := range table {
		for col, colValue := range rowValue {
			if colValue == x {
				fmt.Printf("found %d(row:%d, col:%d)\n", x, row, col)
				continue LOOPCONTINUE
			}
		}
	}
}

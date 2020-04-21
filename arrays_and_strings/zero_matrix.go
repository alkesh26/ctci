package main

import (
	"fmt"
)

func main() {
	array := [][]int8{
			{0, 1, 2, 3},
    		{4, 5, 6, 7},
			{9, 5, 1, 4},
			{4, 7, 12, 7},
			{7, 12, 1, 1},
		}

	array = processArrayZeros(array)

	fmt.Println(array)
}

func processArrayZeros(array [][]int8) [][]int8{
	m, n := 5, 4

	rowHasZero := false
	columnHasZero := false

	for i := 0; i < n; i++ {
		if array[0][i] == 0 {
			rowHasZero = true
			break
		}
	}

	for i := 0; i < m; i++ {
		if array[i][0] == 0 {
			columnHasZero = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if array[i][j] == 0 {
				array[0][j] = 0
				array[i][0] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if array[0][j] == 0 || array[i][0] == 0 {
				array[i][j] = 0
			}
		}
	}

	if rowHasZero {
		for i := 0; i < n; i++ {
			array[0][i] = 0
		}
	}

	if columnHasZero {
		for i := 0; i < m; i++ {
			array[i][0] = 0
		}
	}

	return array
}

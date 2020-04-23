package main

import "fmt"

func main() {
	array := [][]int8{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	N := 4
	array = rotateMatrix(array, N)

	fmt.Println(array)
}

func rotateMatrix(array [][]int8, N int) [][]int8 {
	for step := 0; step < N/2; step++ {
		first, last := step, N-1-step
		for i := first; i < last; i++ {
			offset := i - first
			tmp := array[first][i]
			array[first][i] = array[last-offset][first]
			array[last-offset][first] = array[last][last-offset]
			array[last][last-offset] = array[i][last]
			array[i][last] = tmp
		}
	}
	return array
}

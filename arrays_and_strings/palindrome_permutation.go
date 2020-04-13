package main

import (
	"fmt"
)

func main() {
	x := "mm a ad"
	bit := 0

	for i := 0; i < len(x); i++ {
		if x[i] != ' ' {
			index := x[i] - 'a'

			bit = bit ^ (1 << index)
		}
	}

	fmt.Println(bit&(bit-1) == 0)
}

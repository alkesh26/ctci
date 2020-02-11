package main

import (
	"fmt"
)

func isPalindrome(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	arr := make([]int, 128)

	for i := 0; i < len(a); i++ {
		index := a[i]
		arr[index]++
	}

	for i := 0; i < len(b); i++ {
		index := b[i]
		arr[index]--
		if arr[index] < 0 {
			return false
		}
	}

	return true
}

func main() {
	a, b := "aaa", "aaa"
	fmt.Println(isPalindrome(a, b))
}

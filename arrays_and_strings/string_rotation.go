package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "ixrecreatedmatr"
	s2 := "matrix"

	fmt.Println(strings.Contains(s1+s1, s2))

	s3 := "dasixrecreatedmatr"
	s4 := "matrix"

	fmt.Println(strings.Contains(s3+s3, s4))
}

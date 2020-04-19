package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "aaa"

	if len(str) == 0 {
		fmt.Println("0 length string")
	} else {
		computeCompressedStr(str)
	}

}

func computeCompressedStr(str string) {
	strLength := len(str)

	compressedStr := ""
	count := 0

	for i := 0; i < strLength; i++ {
		count++

		if i + 1 >= strLength || str[i+1] != str[i] {
			compressedStr += string(str[i])
			compressedStr += strconv.Itoa(count)
			count = 0
		}
	}

	compressedStrLength := len(compressedStr)

	if compressedStrLength < strLength {
		fmt.Println(compressedStr)
	} else {
		fmt.Println(str)
	}
}

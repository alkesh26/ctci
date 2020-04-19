package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "aaaaaaaabc"

	if len(str) == 0 {
		fmt.Println("0 length string")
	} else {
		computeCompressedStr(str)
	}

}

func computeCompressedStr(str string) {
	strLength := len(str)

	compressedStr := ""
	i := 0

	for i < strLength-1 {
		count := 1
		for i < strLength-1 && str[i+1] == str[i] {
			count++
			i++
		}

		compressedStr += string(str[i])
		compressedStr += strconv.Itoa(count)
		i++
	}

	if i == strLength-1 {
		compressedStr += string(str[i])
		compressedStr += "1"
	}

	compressedStrLength := len(compressedStr)

	if compressedStrLength < strLength {
		fmt.Println(compressedStr)
	} else {
		fmt.Println(str)
	}
}

package main
import "fmt"

func main() {
	str := "Mr Sam Billings    "
	trueLength := 15

	spaceCount := 0
	for i := 0; i < trueLength; i++ {
		if str[i] == ' '{
			spaceCount++
		}
	}

	index := trueLength + spaceCount*2

	runeString := []rune(str)
	for i := trueLength - 1; i >= 0; i-- {
		if runeString[i] == ' ' {
			runeString[index - 1] = '0'
			runeString[index - 2] = '2'
			runeString[index - 3] = '%'
			index -= 3
		} else {
			runeString[index-1] = runeString[i]
			index--
		}

	}

	fmt.Println(string(runeString))
}

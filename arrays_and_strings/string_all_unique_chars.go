package main
import "fmt"

func main(){
	var str string
	fmt.Scanf("%s", &str)

	checker := 0
	hasUnique := true
	for i := 0; i < len(str); i++ {
		bitIndex := str[i] - 'a'

		if checker & (1 << bitIndex) > 0 {
			hasUnique = false
			break;
		}

		checker = checker | (1 << bitIndex)
	}

	fmt.Println("String has unique characters")
	fmt.Println(hasUnique)
}

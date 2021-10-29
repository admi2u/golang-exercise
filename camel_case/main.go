package main

import "fmt"

func main() {
	camelString := "saveChangesInTheEditor"
	var words []string

	i := 0
	for {
		if i > len(camelString) {
			break
		}

		if i == len(camelString) {
			word := camelString[0:i]
			words = append(words, word)
			break
		} else if camelString[i] >= 'A' && camelString[i] <= 'Z' && i != 0 { // 注意这里使用单引号
			word := camelString[0:i]
			words = append(words, word)
			camelString = camelString[i:]
			i = 0
		} else {
			i++
		}
	}

	fmt.Printf("The camelCase string contains %d words:\n", len(words))
	for _, w := range words {
		fmt.Println(w)
	}
}

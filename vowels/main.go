package main

import (
	"fmt"
	"strings"
)

/*
write a function that returns the number of vowels
used in a string. vowels are the characters 'a' 'e' 'i' 'o' 'u'
*/

func vowels(inString string) int {
	sum := 0
	inString = strings.ToLower(inString)
	for i := 0; i < len(inString); i++ {
		if inString[i] == 'a' || inString[i] == 'e' || inString[i] == 'i' ||
			inString[i] == 'o' || inString[i] == 'u' {
			sum++
		}
	}
	return sum
}

func main() {
	fmt.Println(vowels("why DOO you ask"))
}

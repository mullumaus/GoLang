package main

import (
	"fmt"
	"strings"
)

/*
write a function that accepts a string. THe function should capitalize the first
letter of each word in the string then return the capitalized string
for example:
capitalize('a short sentence') -> "A short sentence"
*/

func capitalize(inStr string) string {
	outString := ""
	wordlist := strings.Split(inStr, " ")
	for _, word := range wordlist {
		if len(word) > 1 {
			word = strings.ToUpper(string(word[0])) + string(word[1:])
		} else {
			word = strings.ToUpper(string(word[0]))
		}
		outString = outString + word + " "
	}
	return outString
}

func capitalize1(inStr string) string {
	outString := strings.ToUpper(string(inStr[0]))
	for i := 1; i < len(inStr); i++ {
		if inStr[i-1] == ' ' {
			outString = outString + strings.ToUpper(string(inStr[i]))
		} else {
			outString = outString + string(inStr[i])
		}
	}
	return outString
}

func main() {
	instr := "Welcome, to the, online portal, of GeeksforGeeks"
	fmt.Println("Before ", instr)
	fmt.Println(capitalize(instr))
	fmt.Println(capitalize1(instr))
}

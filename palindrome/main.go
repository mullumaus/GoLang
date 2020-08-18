package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	if len(str) == 1 {
		return str
	}
	return reverseString(str[1:]) + string(str[0])
}

func palindrom(str string) bool {
	if str == reverseString(str) {
		return true
	}
	return false
}

func palindrom2(str string) bool {
	n := len(str) / 2
	lastChar := len(str) - 1
	for i := 0; i <= n; i++ {
		if str[i] != str[lastChar-i] {
			return false
		}
	}
	return true
}
func main() {
	//str := "02022020"
	//str := "hello there"
	str := "Red rum, sir, is murder"
	str = strings.ToLower(str)
	//palindrom(strings.ToLower(str))
	if palindrom2(str) {
		fmt.Println("The string is palindrome")
	} else {
		fmt.Println("This string is not palindrome")
	}
}

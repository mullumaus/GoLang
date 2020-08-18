package main

import "fmt"

func reverseString(str string) string {
	if len(str) == 1 {
		return str
	}
	return reverseString(str[1:]) + string(str[0])
}

func main() {
	result := reverseString("hello world")
	fmt.Println(result)
}

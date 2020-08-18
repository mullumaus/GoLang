package main

import "fmt"

/* Given an integer, return an integer that is the reverse
for example,
reverseInt(15) = 51
reverseInt(900) = 9
reverseInt(-90) = -9
*/

func reverseInt(num int) int {
	newInt := 0

	for num != 0 {
		remainder := num % 10
		newInt = newInt * 10
		newInt += remainder
		num = num / 10
	}

	return newInt
}
func main() {
	newInt := reverseInt(-1000)
	fmt.Println(newInt)
}

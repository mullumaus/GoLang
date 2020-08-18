package main

import "fmt"

/* Given a string, return the character that is most commonly used in the string*/
func maxchar(str string) {
	charMap := map[byte]int{}
	for i := 0; i < len(str); i++ {
		c := str[i]
		val, ok := charMap[c] //ok=true if key c exist in map
		if ok {
			charMap[c] = val + 1
		} else {
			charMap[c] = 1
		}
	}
	max := 0
	maxchar := byte(0)
	for key, value := range charMap {
		if value > max {
			max = value
			maxchar = key
		}
	}
	fmt.Printf("char %c occurs %d times\n", maxchar, max)
}

func main() {
	maxchar("fashjjjjjsds")
}

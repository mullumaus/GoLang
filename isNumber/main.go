/*
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
*/

package main

import "unicode"

func getCol(c byte) int {
	col := map[string]int{
		"sign":   0,
		"number": 1,
		".":      2,
		"exp":    3,
		"other":  4,
		"blank":  5}
	var key string
	if unicode.IsDigit(rune(c)) {
		key = "number"
	} else if c == '+' || c == '-' {
		key = "sign"
	} else if c == 'e' || c == 'E' {
		key = "exp"
	} else if c == '.' {
		key = "."
	} else if c == ' ' {
		key = "blank"
	}
	return col[key]
}
func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	transTable := [][]int{{1, 2, 7, -1, -1, 0},
		{-1, 2, 7, -1, -1, -1},
		{-1, 2, 3, 4, -1, 9},
		{-1, 3, -1, 4, -1, 9},
		{6, 5, -1, -1, -1, -1},
		{-1, 5, -1, -1, -1, 9},
		{-1, 5, -1, -1, -1, -1},
		{-1, 8, -1, -1, -1, -1},
		{-1, 8, -1, 4, -1, 9},
		{-1, -1, -1, -1, -1, 9}}
	legalState := []int{2, 3, 5, 8, 9}
	state := 0
	for i := 0; i < len(s); i++ {
		col := getCol(s[i])
		state = transTable[state][col]
		if state == -1 {
			return false
		}
	}
	for _, st := range legalState {
		if state == st {
			return true
		}
	}
	return false
}

func main() {

}

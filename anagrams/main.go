package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type sortString []byte

func (s sortString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortString) Len() int {
	return len(s)
}

func (s sortString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func trimString(instr string) string {
	outstring := ""
	instr = strings.ToLower(instr)
	for _, r := range instr {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			outstring += string(r)
		}
	}
	return outstring
}

func anagrams(instr1 string, instr2 string) bool {
	s1 := sortString(trimString(instr1))
	s2 := sortString(trimString(instr2))
	sort.Sort(s1)
	sort.Sort(s2)

	fmt.Println(string(s1))
	fmt.Println(string(s2))
	if string(s1) == string(s2) {
		return true
	}
	return false
}

func main() {
	s1 := "Rail safety"
	s2 := "fairy tales"
	/* 	s1 := "Hi there"
	   	s2 := "Bye there" */
	if anagrams(s1, s2) {
		fmt.Println(s1, "and", s2, "is anagrams")
	} else {
		fmt.Println(s1, "and", s2, "is NOT anagrams")
	}
}

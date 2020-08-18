package main

import "fmt"

/*
write a function that accepts a positive number N
The function should console log a step shape with N
levels using the #character. Make sure the step has spaces
on the right hand side
for example:
steps(2):
	'# '
	'##'
step3(3)
	'#  '
	'## '
	'###'
*/

func steps(n int) {
	str := ""
	fmt.Println(n)
	for i := 1; i <= n; i++ {
		str = str + "#"
		padded := ""
		if n-len(str) > 0 {
			padded = fmt.Sprintf("%*s", n-len(str), " ")
		}
		result := str + padded
		fmt.Println(result)
	}
}

func steps1(n int) {
	fmt.Println(n)
	for i := 1; i <= n; i++ {
		result := ""
		for j := 1; j <= n; j++ {
			if j <= i {
				result += "#"
			} else {
				result += " "
			}
		}

		fmt.Println(result, "|")
	}
}

func printNum(n int, row int, str string) {
	if n == row {
		return
	}
	if n == len(str) {
		fmt.Println(str)
		printNum(n, row+1, "")
		return
	}

	if len(str) <= row {
		str += "#"
	} else {
		str += " "
	}
	printNum(n, row, str)

}

func main() {
	//steps(10)
	//steps1(10)
	printNum(10, 0, "")
}

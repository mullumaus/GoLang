package main

import "fmt"

/*
write a function that accepts a positive number N
the function should console log a pyramid shape
with N levels using the # character. Make sure the
pyramid has space on both the left and right hand sides
pyramid(1)
	'#'
pyramid(2)
	" # "
	"###"
pyramind(3)
	"  #  "
	" ### "
	"#####"
pyramind(4)
	"   #  "
	"  ###  "
	" ##### "
	"#######"
*/

func getcolumn(n int) int {
	return 2*n - 1
}

func pyramid(n int) {
	for row := 0; row < n; row++ {
		midpoint := (2*n - 1) / 2
		level := ""
		for column := 0; column < 2*n-1; column++ {
			if midpoint-row <= column && midpoint+row >= column {
				level += "#"
			} else {
				level += " "
			}
		}
		fmt.Println(level)
	}

}

func pyramid1(n int, row int, level string) {
	column := 2*n - 1
	if row == n {
		return
	}
	if len(level) == column {
		fmt.Println(level)
		pyramid1(n, row+1, "")
		return
	}
	midpoint := (2*n - 1) / 2
	addstr := ""
	if midpoint-row <= len(level) && midpoint+row >= len(level) {
		addstr = "#"
	} else {
		addstr = " "
	}
	pyramid1(n, row, level+addstr)
}

func main() {
	n := 10
	//pyramid(n)
	pyramid1(n, 0, "")
}

/*
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
*/

package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	res := []int{}
	for i := 1; i < int(math.Pow(10, float64(n))); i++ {
		res = append(res, i)
	}
	return res
}

func main() {
	fmt.Println(printNumbers(5))
}

/*
实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
*/
package main

import "fmt"

func myPow(x int, n int) int {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	if n%2 == 1 {
		return x * myPow(x*x, n/2)
	}
	return myPow(x*x, n/2)

}
func main() {
	fmt.Println(myPow(3, 5))
}

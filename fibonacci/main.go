package main

import "fmt"

func main() {
	n := 20
	fibList := make([]int64, n+1)
	for i := 0; i <= n; i++ {
		fib2(i, fibList)
	}
	for i := 0; i <= n; i++ {
		fmt.Println(fibList[i])
	}

}

/* func fib(n int) int64 {
	if n < 2 {
		return int64(n)
	}
	return fib(n-1) + fib(n-2)
} */

func fib2(n int, fblist []int64) {
	if n < 2 {
		fblist[n] = int64(n)
		return
	}
	fblist[n] = fblist[n-1] + fblist[n-2]
}

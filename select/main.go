package main

import "fmt"

//The select statement lets a goroutine wait on multiple
//communication operations. A select blocks until one of its cases can run, then it executes that case.
//It chooses one at random if multiple are ready.

func fabonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func main() {
	c := make(chan int, 1)
	quit := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fabonacci(c, quit)
}

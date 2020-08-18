package main

import "fmt"

/* func main() {
	c := make(chan string)
	c <- "hello there!" //block here, By default, sends and receives block until the other side is ready.
	fmt.Println(<-c)
} */

func main() {
	c := make(chan string)
	for i := 0; i < 4; i++ {
		go printString("Hello there!", c)
	}

	for s := range c {
		fmt.Println(s)
	}
}

func printString(s string, c chan string) {
	fmt.Println(s)
	c <- "Done printing."
}

/* func main() {
	c := make(chan string)

	for i := 0; i < 4; i++ {
		go printString("Hello there!", c)
	}

	for {
		fmt.Println(<-c)
	}
}

func printString(s string, c chan string) {
	fmt.Println(s)
	c <- "Done printing."
} */

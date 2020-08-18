package main

import (
	"fmt"
	"runtime"
	"time"
)

//switch without a condition is the same as switch true
func noCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os { //short statement in switch
	case "darwin":
		fmt.Println("OS x.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	noCondition()
}

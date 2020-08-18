package main

import (
	"log"
	"time"
)

func main() {
	//1 mean buffer size,
	//0  mean Un-buffered channels are only writable
	//when there's someone blocking to read from it
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	select {
	case <-ch:
		log.Println("read")
	case <-timeout:
		log.Println("read from ch has timeout")
	}
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		//fmt.Println(link)
		go checkLink(link, c) //run this function in a new go routine

	}

	/* for i := 0; i < len(links); i++ {
		go checkLink(<-c, c)
	} */

	/* 	for {
		go checkLink(<-c, c)
		time.Sleep(10 * time.Second)
	} */
	for l := range c { //wait for channel to return something
		go func(link string) { //literal function or lambda
			checkLink(link, c)
			time.Sleep(10 * time.Second)
		}(l) //the parameter pass to literal
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- "Might be down I think"
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWrite struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Eror:", err)
		os.Exit(1)
	}
	lw := logWrite{}

	/* 	bs := make([]byte, 999999) //slice with 999999 empty elements
	   	resp.Body.Read(bs)
		   fmt.Println(string(bs)) */
	io.Copy(lw, resp.Body)
}

func (logWrite) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

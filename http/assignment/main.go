package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File enter the file name to be read")
		os.Exit(1)
	}
	fileName := os.Args[1]
	fd, err := os.Open(fileName)
	if err != nil {
		fmt.Println("failed to open file ", fileName)
		os.Exit(1)
	}
	io.Copy(os.Stdout, fd)
}

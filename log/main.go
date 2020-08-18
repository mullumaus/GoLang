package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("filename.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

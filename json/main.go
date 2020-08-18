package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		panic("Failed to run jason encoder")
	}
	fmt.Println(b)

	//decoder
	var msg message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		panic("Failed to run json decoder")
	}
	fmt.Println(msg.Name)
}

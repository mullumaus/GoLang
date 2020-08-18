package main

import "fmt"

func main() {
	//create empty map, three method to create map
	/*var colors map[string]string
	colors := make(map[string]string)
	colors["red"] = "#FF0000"
	*/
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#FF0001",
	}

	//delete map key
	/* delete(colors, "red") */

	//iterate over map

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for color ", color, " is ", hex)
	}
}

package main

import "log"

//Gift is
type Gift struct {
	sender    string
	recipient string
	number    int
	contents  string
}

type reindeer string

func (r reindeer) takeoff() {
	log.Printf("%q lifts off", r)
}

func (r reindeer) land() {
	log.Printf("%q gently lands", r)
}

func (r reindeer) togglenose() {
	if r != "rudolph" {
		panic("Invalid reindeer operation")
	}
	log.Printf("%q node changes state", r)
}
func main() {

}

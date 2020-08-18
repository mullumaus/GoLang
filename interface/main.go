package main

import "fmt"

//Dog is
type Dog struct {
	Age interface{} //empty interface type, you can assign it a value of any type
}

func emptyInterface() {
	dog := Dog{}
	dog.Age = "3"
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)

	dog.Age = 10
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)

	dog.Age = "not really an age"
	fmt.Printf("%#v %T\n", dog.Age, dog.Age)
}

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	emptyInterface()
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

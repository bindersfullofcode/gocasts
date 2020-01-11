package main

import "fmt"

type Greeter interface {
	Greet() string
}

func PrintGreeting(greeter Greeter) {
	fmt.Println(greeter.Greet())
}

type GreetingString string

func (g GreetingString) Greet() string {
	return fmt.Sprintf("Hello, %v!", g)
}

type GreetingPerson string

func (g)

func main() {
	greeting := GreetingString("Cooper")
	PrintGreeting(greeting)
}

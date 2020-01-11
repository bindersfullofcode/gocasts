package main

import "fmt"

type Employee struct {
	Name  string
	Title Title
}

type Title string

func (t Title) Print() {
	fmt.Printf("Title=%v\n", t)
}

func main() {
	cooper := Employee{
		Name:  "cooper",
		Title: "Software Developer",
	}

	fmt.Printf("%+v\n", cooper)
	cooper.Title.Print()
}

package main

import "fmt"

type Animal interface {
	Says() string
	NumLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberTeeth int

}

func main() {

	dog := Dog {
		Name: "Lilly",
		Breed: "Poodle",
	}
	PrintInfo(&dog)

	gorilla := Gorilla {
		Name: "Big Boy",
		Color: "black",
		NumberTeeth: 40,
	}
	PrintInfo(&gorilla)


}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumLegs(),"legs")
}

func (g *Gorilla) Says() string {
	return "ahhh"
}
func (g *Gorilla) NumLegs() int {
	return 4 
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumLegs() int {
	return 4
}
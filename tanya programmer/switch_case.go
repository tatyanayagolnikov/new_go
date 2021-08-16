package main

import "fmt"

func main() {

	newVar := "TATBOT PROGRAMMER"

	switch newVar {
	case "TATBOT":
		fmt.Println("#1")
	case "TATBOT PROGRAMMER": // This works
		fmt.Println("TATBOT The programmer")
	default: 
	    fmt.Println("None of these work")
	}



	myVar := "snail"

	switch myVar {
	case "cat":
		fmt.Println("myVar is a cat")
	case "dog":
		fmt.Println("myVar is a dog")
	case "snail":   // This works 
		fmt.Println("myVar is a snail")
	default:
		fmt.Println("none of the above")
	}




	fmt.Println("Tanya Programmer")
}
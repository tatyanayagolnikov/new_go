package main

import "fmt"

func main() {
	myNum := 42
	isTrue := false

	if myNum > 41 && isTrue {
		fmt.Println("woohoo")
	} else if myNum <100 && isTrue {
		fmt.Println("dope")
	} else if myNum == 99 || !isTrue {
		fmt.Println("and || works")
	}


	fmt.Println("Tanya Programmer")

}
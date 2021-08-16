package main

import "fmt"

// -------- MAP SYNTAX---------
//    myMap := make(map[string]string) 

type User struct {
	First string
	Last string
}

type Person struct {
	Sport string
	Food string
}




func main() {
	
	myMap := make(map[string]User)
	newMap := make(map[string]string)
	newMap2 := make(map[int]string)
	newMap3 := make(map[string]Person)
	newMap4 := make(map[int]Person)
	newMap5 := make(map[string]User)

	p3 := User {
		First: "Little T",
		Last: "dast PROGRAMMER",
	}

	newMap5["GOO"] = p3
	fmt.Println(newMap5["GOO"].First, newMap5["GOO"].Last)

	p2 := Person {
		Sport: "skiing",
		Food: "shrimp",
	}

	newMap4[2] = p2
	newMap4[5] = p2
	fmt.Println(newMap4[2].Sport)
	fmt.Println(newMap4[5].Food)


	p := Person {
		Sport: "Basketball",
		Food: "Sushi",
	}
	newMap3["p"] = p
	fmt.Println(newMap3["p"].Sport)

	newMap["What"] = "Rad!"
	fmt.Println(newMap["What"]) 

	newMap2[1] = "buns"
	newMap2[2] = "funs"
	fmt.Println(newMap2[1])
	fmt.Println(newMap2[2])

	u := User {
		First: "Tanya",
		Last: "Programmer",
	}
	myMap["u"] = u
	fmt.Println(myMap["u"].First)
	fmt.Println(myMap["u"].Last)
	
}
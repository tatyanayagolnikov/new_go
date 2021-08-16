package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9} // COMPOSITE LITERAL []int

	names := []string{"Little T","Big B","Smith"}
	fmt.Println(names[0:2])

	fmt.Println(numbers[4:8]) //



	var mySlice []string // This is a SLICE of String

	mySlice = append(mySlice, "Tanya") // APPEND (adding) to a slice
	mySlice = append(mySlice, "Programmer")
	mySlice = append(mySlice, "is rad!")

	sort.Strings(mySlice)
	
	fmt.Println(mySlice)


}
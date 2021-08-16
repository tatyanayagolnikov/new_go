package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	} 
	fmt.Println("Tanya Programmer")


	animals := []string{"cat","dog","frog","hog","snail"}
	for i, v := range animals {
		fmt.Println(i, v)
	} 

	sports := make(map[string]string)
	sports["first"] = "bball"
	sports["second"] = "skiing"
	sports["third"] = "skate"
	for i, v := range sports {
		fmt.Println(i, v)
	}

	type User struct {
		First string
		Last string
		Age int
		Food string
	}

	var users []User
	users = append(users, User{"Tat","Programmer",35, "bball"})
	users = append(users, User{"Big", "Daddy", 42, "skiing"})
	users = append(users, User{"Snow", "White", 17, "singing"})
	
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age, v.Food)
	}

	

}
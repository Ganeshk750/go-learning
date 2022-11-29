package main

import "fmt"

func main(){
	var name string
	name = "Golang"
	// Condition
	if name != "Golang"{
		fmt.Println("Matched Name:", name)
	}
	fmt.Println("Defalut")

	// If else 

    age := 18
	if age == 18 {
		fmt.Println("You are eligible for vote")
	} else{
		fmt.Println("You are not eligible for vote")
	}

	if false{
		fmt.Println("Hey")
	}else{
		fmt.Println("Hello")
	}
}

package main

import "fmt"

func main(){
	fmt.Println("Hey! Welcome To Go World!")
	str := firstFunc()
	fmt.Println(str)
	out := add(5, 5)
	fmt.Println("Addition Output=", out)
	wish := toHelloWithName("Golang")
	fmt.Println(wish)
}

func firstFunc() string{
	return "My Name is Gooo"
}

func add(x int, y int) int{
	sum := 0
	sum = x + y
	return sum
}

func toHelloWithName(name string) string{
	return "Hello !" + name
}


package main

import "fmt"

func main() {
	//pointer declaration
	var x *int
	y := 10
	x = &y

	fmt.Println(y, "->", x)
	fmt.Println(*x, "->", x)
	fmt.Println(y, "->", &y)
	fmt.Println(*x, "->", &x)
}
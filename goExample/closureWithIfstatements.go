package main

import "fmt"

func main() {

	var r int
	fmt.Println("enter a number:")
	fmt.Scanf("%d", &r)
	//a closure : its an anonymous function
	x := func(r int) int {
		return r % 2
	}(r) // imp. this makes a closure special
	if x == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
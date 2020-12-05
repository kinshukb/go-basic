package main

import (
	"fmt"
)

//func main() {
	//scope of variable changes
	x := 9
	for x := 1; x < 15; x++ { //here another declaration of x
		fmt.Println(x)
	}
	fmt.Println(x)

	//different scope of x
	/*x := 9
	for x = 1; x < 15; x++ {
		fmt.Println(x)
	}*/
	fmt.Println(x)
	// fmt.Println("enter a value:")
	// fmt.Scanf(y)

}

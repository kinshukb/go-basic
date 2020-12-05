package main

import "fmt"

func main() {

	//i := ""                   //no dead-code allowed in golang...wow
	j := 0
	fmt.Println("enter choice:")
	fmt.Println("***")
	//fmt.Scanf("%s", &i)              //scanf works just like c lang
	fmt.Scanf("%d", &j) //user input
	fmt.Println("**************")
	switch j {
	case 1:
		fmt.Println("hii 1")
	case 2:
		fmt.Println("hii 2")
	case 3:
		fmt.Println("hii 3")
	default:
		fmt.Println("hii default")

	}
}
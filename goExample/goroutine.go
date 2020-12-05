package main

import "fmt"

func addDigit(x int) {
	var newNo int
	if x != 0 {
		rem := x % 10
		num := x / 10
		newNo := rem + num

		//	addDigit(newNo)
		//	return newNo
	}
	fmt.Println(newNo)
	//return newNo
}
func main() {
	x := 54
	go addDigit(x)
	//fmt.Println(res)

}
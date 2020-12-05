package main

import "fmt"

func main() {
	m2 := make(map[int]int, 20)
	m2[1] = 1
	fmt.Println(len(m2))
	for j := 0; j < len(m2); j++ {
		m2[j] = j
		fmt.Println("*******")
		fmt.Println(m2[j])
	}
	fmt.Println(m2)
}
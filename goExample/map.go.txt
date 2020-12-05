package main

import "fmt"

func main() {

	//maps are associative : for retrieval in goes the KEY, out comes the VALUE ,k:v pair
	//maps are unordered
	//declaring map -using make func
	m1 := make(map[int]int)
	for i := 0; i < 10; i++ {
		m1[i] = i
		fmt.Println(m1[i])
	}
	m2 := make(map[int]int, 20)
	for j := 0; j < len(m2); j++ {
		m2[j] = j
		fmt.Println("*******")
		fmt.Println(m2[j])
	}
	//deletes corr. map element with corr. key
	delete(m1, 1)
	fmt.Println("length of map:", len(m1))
	//kinda toString()
	fmt.Println(m1) //state
	m1[1] = 1
	fmt.Println("length of map:", len(m1))
	fmt.Println(m1)

	//2nd way of declaring and initializing map
	//without using make
	m3 := map[string]int{"k1": 10, "k2": 11}
	fmt.Println(m3)
}
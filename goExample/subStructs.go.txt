package main

import (
	"fmt"
)

//user defined struct declaration
type car struct {
	category string
	rating   int
	name     string
}

//a member func of car struct binding done to conceptualize class
// GetDetails
func (c car) GetDetails() {
	fmt.Println(c.name)
	fmt.Println(c.rating, c.category)
}

type ferrari struct {
	car
	topSpeed float64
	color    string
}

// a member function of ferrari struct binding done to conceptualize class
func (f ferrari) yeahRed() {
	fmt.Println("Top Speed:", f.topSpeed, "Color:", f.color)
	fmt.Println("Red ferrari...thats some color")
}

type astonMartin struct {
	car
	topSpeed float64
	color    string
}

// a member function of astonMartin struct binding done to conceptualize class
//this is how it happens in golang
func (a astonMartin) bond() {
	fmt.Println("Top Speed:", a.topSpeed, "Color:", a.color)
	fmt.Println("HI Bond, James Bond!")
}

//interface declaration
type Icar interface {
	engine() int
	wheels() int
	isElectric() bool
}

//very important to achieve loose coupling
func intfFunc(i Icar) {
	fmt.Println(i.engine())
	fmt.Println(i.isElectric())
	fmt.Println(i.wheels())
}

//ferrari's implementation
func (f ferrari) engine() int {
	fmt.Println("Great Engine!!")
	return 4
}
func (f ferrari) isElectric() bool {
	fmt.Println("This Ferrari is not electric!")
	return false
}
func (f ferrari) wheels() int {
	fmt.Println("Blah! It has 4 wheels like any car")
	return 4
}

//aston Martin's implementation
func (a astonMartin) engine() int {
	fmt.Println("Great Engine, Mr Bond!!")
	return 4
}
func (a astonMartin) isElectric() bool {
	fmt.Println("This aston martin is not electric!")
	return false
}
func (a astonMartin) wheels() int {
	fmt.Println("Blah! It has 4 wheels like any car")
	return 4
}
func main() {
	//car1 := car{"A Class", 4}
	car2 := car{"S Class", 5, "Ferrari"}
	//initializing while declaring kinda calling constructor
	aston := astonMartin{car{"A Class", 4, "AstonMartin"}, 190, "Silver"}
	ferrari := ferrari{car2, 220, "Red"}
	aston.getDetails()
	aston.bond()
	ferrari.getDetails()
	ferrari.yeahRed()
	intfFunc(ferrari) //upcasting happening during fn call
	intfFunc(aston)   //upcasting happening during fn call
}
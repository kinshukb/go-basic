package main

import (
	"fmt"
	"math"
)

type rect struct {
	leng float64
	bred float64
}
type circle struct {
	rad float64
}
type areaIntf interface {
	area() float64
}

func (r rect) area() float64 {
	return r.bred * r.leng
}
func (c circle) area() float64 {
	return c.rad * c.rad * math.Pi
}
func calculate(a areaIntf) {
	fmt.Println(a)
	fmt.Println(a.area())
}
//func main() {
	var x int
	var l, b float64
	var r rect
	var radius float64
	var c circle
	fmt.Println("enter 1 for rect and 2 for circle:")
	fmt.Scanf("%d", &x)
	if x == 1 {
		fmt.Println("enter length and breadth:")
		fmt.Scanf("%f", &l)
		//	fmt.Println("******")
		fmt.Scanf("%f", &b)
		r.leng = l
		r.bred = b
		//float64 result=area(r)
		calculate(r)

	} else {
		fmt.Println("enter radius:")
		fmt.Scanf("%f", &radius)
		c.rad = radius
		//var res=area(c)
		//fmt.Println(res)
		calculate(c)

	}
}

/*var x=10
if x>5{
	fmt.Println("greater than")
}
else{
	fmt.Println("small...")
}*/
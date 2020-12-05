package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}
type emp struct {
	person
	sal   float64
	empno int
}

func (e emp) getSal() float64 {
	return e.sal
}

//imp for loose coupling
func myIntfFunc(s salary) {
	fmt.Println("*****", s.getSal())
}

type salary interface {
	getSal() float64
	getIncrement() float64
}

func (e emp) getIncrement() float64 {
	return 2323.2323
}
//func main() {
	p1 := person{"dhsdh", "ywyewe"}
	emp1 := emp{p1, 446545, 1}
	x := emp1.getSal()
	y := emp1.getIncrement()
	fmt.Println(x)
	fmt.Println(y)
	myIntfFunc(emp1) //so here it checks whether emp1 belonging to
	//emp struct has implemented interface fully or not when this function is called eventually upcasting happens
}

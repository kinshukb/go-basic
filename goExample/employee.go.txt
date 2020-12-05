package main

import (
	"fmt"
)

type owner struct {
	name         string
	sharePercent float64
	address      string
}
type emp struct {
	name    string
	empno   int
	address string
}
type salary interface {
	getSal() float64
}

func (e emp) getSal() float64 {
	if e.empno > 30 {
		return 8000.55
	} else {
		return 4000
	}
}
func (o owner) getSal() float64 {
	if o.sharePercent > 20 && o.sharePercent < 40 {
		return 5422354
	} else if o.sharePercent > 40 {
		return 87786454
	} else {
		return 564888
	}
}

func main() {

	emp1 := emp{"king", 5, "pune"}
	own1 := owner{"queen", 30, "mumbai"}
	sal1 := emp1.getSal()
	sal2 := own1.getSal()
	fmt.Println(sal1)
	fmt.Println(sal2)
}
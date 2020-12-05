package main

import (
	"fmt"
)

type jedi interface {
	hasForce() bool
}

type obiWan struct {
	num int
}

func (obi *obiWan) hasForce(ch chan<- bool) {
	ch <- true
}
func (obi *obiWan) hasForce() {
	fmt.Println("may the force be with you!!")
}
func main() {

	// obi := obiWan{1}
	// fmt.Println(obi.num)

	// ch := make(chan bool)
	// mrObi := obiWan{1}
	// go mrObi.hasForce(ch)
	// fmt.Println(<-ch)
	// fmt.Println(mrObi.num)

	obi := obiWan{4}

	obi.hasForce()

}
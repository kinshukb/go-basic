package main

import "fmt"

func (o order) doSome(s string) {
	x := make([]int, 4)
	for i := 0; i < 4; i++ {
		x[i] = i
		fmt.Println(x[i])
		//	x[i]

	}

}

type order struct {
	oName string
	id    int
	qty   int
}

type customer struct {
	cName   string
	address string
}

func main() {

	ord := order{"fdf", 4, 5}
	ord.doSome("ghgh")
	cust := new(customer)
	cust.address = "ddd"
	cust.cName = "aaa"
	fmt.Println(cust)

	cust1 := make(customer{"dsd", "ffd"})
	fmt.Println(cust1)

}
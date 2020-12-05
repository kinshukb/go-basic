package main

import (
	"fmt"
)

func check(done chan int) {
	fmt.Println("hiiiiiiii")
	done <- 545
}
func mul(n int, done chan<- int) {
	//arrow required <- for catching or returning any value from a go routine
	result := n * 5
	done <- result
}
func printStr(s string, newChan chan<- string) {
	newChan <- s
}

func main() {
	//1st channel declaration
	done := make(chan int)

	//starting new path of execution
	go mul(5, done)
	fmt.Println(<-done) //expecting 25

	//to be performed by main thread
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}

	//starting new path of execution
	go check(done)
	//<-done
	fmt.Println(<-done) //545

	fmt.Println("in main....")

	//2nd channel
	newChan := make(chan string)
	//starting new path of execution

	go printStr("gogogog", newChan)
	fmt.Println(<-newChan)
	//since channels are used we can avoid time.sleep as the main thread will now communicate
	//with go routines using channels

	//time.Sleep(time.Millisecond)

}
package main

import (
	"fmt"
	"time"
)

func testConcurncy(s string) string {
	return s + "blah.."
}

func testConcurncy2(s string, ch chan<- string) {
	s = s + "modifyStr"
	ch <- s
}
func testBuffChan(ch chan<- string) {
	ch <- "empty"

}
func main() {
	//make is used to create a data struct like channel,slice and map
	ch := make(chan string)
	go testConcurncy("blah")
	time.Sleep(time.Millisecond)
	//str := testConcurncy("blah")
	//fmt.Println(str)
	go testConcurncy2("anything", ch)
	fmt.Println(<-ch)

	// a buffered channel
	ch2 := make(chan string, 3)
	ch2 <- "sffdf"
	ch2 <- "klfdfn"
	ch2 <- "jkls"
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	//A buffered channel allows the goroutine that is adding data to the buffered channel to keep running
	//and doing things, even if the goroutines reading from the channel are starting to fall behind a little bit
	go testBuffChan(ch2)
	fmt.Println(<-ch2)
}

  
package main

import (
	"fmt"
)

func sayHello(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {
	ch:=chan int
	
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	sayHello("normalCall")
	go sayHello("concurrentCall")
	//time.Sleep(time.Millisecond)
	sayHello("checCall")
	go sayHello("newCall")
	fmt.Scanln()
	//time.Sleep(time.Millisecond)
}
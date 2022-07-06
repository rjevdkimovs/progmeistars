package main

import (
	"fmt"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("send", i, "to channel")
	}
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	for  {
		fmt.Println(<-ch, "received")
	}	
}



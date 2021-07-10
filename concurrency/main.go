package main

import (
	"fmt"
	"time"
)

func main() {
	var firsttosecond = make(chan int)
	var secondtothirdchannel = make(chan int)
	var endchannel = make(chan int)
	go dotask1(firsttosecond)
	go dotask2(firsttosecond, secondtothirdchannel)
	go dotask3(secondtothirdchannel, endchannel)
	//time.Sleep(2000)
	endingvalue := <- endchannel
	fmt.Printf("End: %v\n",endingvalue)
}

func dotask1(c chan int) {
	c <- 2
	fmt.Printf("task 1 is being done\n")
}

func dotask2(c chan int, c2 chan int) {
	value := <- c
	fmt.Printf("task 2 is being done val : %v\n", value)
	c2 <- value * value
}

func dotask3(c chan int, end chan int) {
	value := <- c
	fmt.Printf("task 3 is being done val : %v\n", value)
	time.Sleep(5000 * time.Millisecond)
	end <- value
}

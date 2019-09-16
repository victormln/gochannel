package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		defer close(ch)
		ch <- "Hola Channel"
	}()
	fmt.Println(<-ch)

	channel1 := make(chan int)
	go func() {
		defer close(channel1)
		channel1 <- 1
		channel1 <- 2
		channel1 <- 3
		channel1 <- 2
		channel1 <- 5
	}()

	for number := range channel1 {
		fmt.Printf("%d", number)
	}

	// Max length is 2
	channel2 := make(chan int, 2)
	channel2 <- 1
	channel2 <- 2
	fmt.Println("\n=========")
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	channel2 <- 3
	fmt.Println("After clean channel:")
	fmt.Println(<-channel2)
}

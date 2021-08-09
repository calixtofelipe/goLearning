package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	for i := 1; i <= 10; i++ {
		go worker(channel)
	}

	for i := 0; i <= 100; i++ {
		channel <- i
	}

}

func worker(channel chan int) {
	for i := range channel {
		fmt.Println(i)
		time.Sleep(3 * time.Second)
	}

}

func forsample() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		time.Sleep(3 * time.Second)
	}
}

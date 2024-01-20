package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func randomGenerator(ch1 chan<- int, ch2 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup when functionda chiqsa

	for i := 0; i < 10; i++ {
		randomNum := rand.Intn(21) + 10 // Random son 10 dan 30 gacha
		ch1 <- randomNum

		if randomNum >= 5 && randomNum <= 10 {
			fmt.Printf("Random son %d, 5 va 10 oralig'ida!\n", randomNum)
		}

		ch2 <- randomNum * randomNum
	}

	close(ch1)
	close(ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go randomGenerator(ch1, ch2, &wg)

	go func() {
		for num := range ch1 {
			fmt.Printf("Received from ch1: %d\n", num)
		}
	}()

	go func() {
		for squaredNum := range ch2 {
			fmt.Printf("Kvadrati: %d\n", squaredNum)
		}
	}()

	wg.Wait()
}

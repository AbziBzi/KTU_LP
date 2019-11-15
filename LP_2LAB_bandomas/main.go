package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Lygiagretaus programavimo 2 laboratorinio darbo bandomasis kolis")

	var wg sync.WaitGroup

	// making all chanels
	sender := make(chan int)
	resultFirst := make(chan int)
	resultSecond := make(chan int)
	isFinished := make(chan bool, 2)

	go sendNumber(0, sender, isFinished)
	go sendNumber(11, sender, isFinished)

	go receiver(sender, resultFirst, resultSecond, &wg, isFinished)

	wg.Add(2)
	go printResults(resultFirst, &wg)
	go printResults(resultSecond, &wg)

	wg.Wait()
	fmt.Println("Done")
}

func sendNumber(from int, sender chan<- int, isFinished chan bool) {
	for i := from; ; i++ {
		select {
		case <-isFinished:
			return
		default:
			sender <- i
		}
	}
}

func receiver(sender <-chan int, resultFirst chan<- int, resultSecond chan<- int, wg *sync.WaitGroup, isFinished chan bool) {
	defer close(resultFirst)
	defer close(resultSecond)
	var count = 0
	for {
		if count < 20 {
			select {
			case i := <-sender:
				if i%2 == 0 {
					resultFirst <- i
				} else {
					resultSecond <- i
				}
			}
			count++
		} else {
			isFinished <- true
			isFinished <- true
			return
		}
	}
}

func printResults(result <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for number := range result {
		fmt.Println(number)
	}
}

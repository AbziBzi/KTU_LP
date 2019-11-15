package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// creating channels that sends data of type int
	sender := make(chan int)
	resultFirst := make(chan int)
	resultSecond := make(chan int)
	// creating signal channel that will say to end infinitive loop in sendNumber func
	isFinished := make(chan bool)

	// Two processes sends data by sender channel
	go sendNumber(0, sender, isFinished)
	go sendNumber(11, sender, isFinished)

	// One process receive data from sender and sends it to result chanel
	// after filter given numbers
	go receiver(sender, resultFirst, resultSecond, isFinished)

	wg.Add(2)
	go printResults(resultFirst, &wg)
	go printResults(resultSecond, &wg)

	wg.Wait()
	fmt.Println("Done")
}

func sendNumber(from int, sender chan<- int, isFinished chan bool) {
	for i := from; ; i++ {
		var message = "sended number: "
		select {
		case <-isFinished:
			return
		case sender <- i:
			fmt.Println(message, i)
		}
	}
}

func receiver(sender <-chan int, resultFirst chan<- int, resultSecond chan<- int, isFinished chan bool) {
	defer close(resultFirst)
	defer close(resultSecond)
	var count = 0
	for {
		if count < 10 {
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

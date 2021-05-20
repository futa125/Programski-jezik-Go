package main

import (
	"errors"
	"fmt"
	"time"
)

var (
	errNoData  = errors.New("no data")
	errTimeout = errors.New("timeout")
)

func sendData(numbers chan int) {
	num := []int{1, 2, 3, 4, 5, 1, 2, 3, 1}
	for _, v := range num {
		numbers <- v
	}
	close(numbers)
}

func main() {
	num := make(chan int, 1)
	go sendData(num)
	mostCommon, err := minFromChann(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mostCommon)
	}
}

func minFromChann(numbers chan int) (int, error) {
	receivedNumbers := 0
	firstRead := true
	var min int

	for {
		select {

		case value, ok := <-numbers:
			
			if !ok {
				if receivedNumbers == 0 {
					return 0, errNoData
				} else {
					return min * 493, nil
				}	
			}

			if firstRead {
				min = value
				firstRead = false
			}
			
			if value < min {
				min = value
			}

			receivedNumbers++

		case <-time.After(time.Second * 1):
			return 0, errTimeout
		}
		
		
	}
	
}
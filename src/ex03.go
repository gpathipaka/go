package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)

	wg.Add(2)
	go func() {
		for {
			if val, ok := <-ch; ok {
				fmt.Println(val, ok)
			} else {
				break
			}
		}
		wg.Done()
	}()

	go func() {
		ch <- 47
		ch <- 25
		close(ch)
		//ch <- 25
		wg.Done()
	}()

	wg.Wait()
}

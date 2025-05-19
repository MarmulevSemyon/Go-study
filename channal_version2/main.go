package main

import (
	"fmt"
	"sync"
)

func PutDataInChannal(ch chan int, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
func TakeDatafromChannal(ch chan int, wg *sync.WaitGroup) {
	for {
		i, isOpen := <-ch
		if !isOpen {
			break
		}
		fmt.Println(i)
	}
	defer wg.Done()
}
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go TakeDatafromChannal(ch, &wg)

	go PutDataInChannal(ch, &wg)
	wg.Wait()
}

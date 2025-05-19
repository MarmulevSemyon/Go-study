package main

import (
	"fmt"
	"sync"
)

func FromArrToChan1(ch1 chan int, arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch1)
	for _, v := range arr {
		ch1 <- v
	}
}
func FromChan1ToChan2(ch1 chan int, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch2)
	for {
		i, isOpenCh1 := <-ch1
		if !isOpenCh1 {
			break
		}
		ch2 <- i * i
	}
}
func FromChan2ToStd(ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		i, isOpenCh2 := <-ch2
		if !isOpenCh2 {
			break
		}
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	ch1, ch2 := make(chan int), make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	wg.Add(3)
	go FromArrToChan1(ch1, arr, &wg)
	go FromChan1ToChan2(ch1, ch2, &wg)
	go FromChan2ToStd(ch2, &wg)
	wg.Wait()
}

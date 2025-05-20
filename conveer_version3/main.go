package main

import (
	"fmt"
)

func FromArrToChan1(ch1 chan int, arr []int, donech chan struct{}) {
	go func() {
		defer close(ch1)
		for _, v := range arr {
			select {
			case ch1 <- v:
			case <-donech:
				return
			}
		}
	}()
}
func FromChan1ToChan2(ch1 chan int, ch2 chan int, donech chan struct{}) {
	defer close(ch2)
	for v := range ch1 {
		select {
		case ch2 <- v * v:
		case <-donech:
			return
		}
	}

}

func main() {
	ch1, ch2, donech := make(chan int), make(chan int), make(chan struct{})
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	defer close(donech)
	go FromArrToChan1(ch1, arr, donech)
	go FromChan1ToChan2(ch1, ch2, donech)
	for v := range ch2 {
		fmt.Println(v)
	}

}

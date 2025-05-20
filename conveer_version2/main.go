package main

import (
	"fmt"
)

func FromArrToChan1(arr []int) chan int {
	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		for _, v := range arr {
			ch1 <- v
		}
	}()
	return ch1
}
func FromChan1ToChan2(ch1 chan int) <-chan int {
	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		for i := range ch1 {
			ch2 <- i * i
		}
	}()
	return ch2
}

func main() {
	// arr := make([]int, 1000)
	// for i := range arr {
	// 	arr[i] = i + 1
	// }
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := FromArrToChan1(arr)
	ch2 := FromChan1ToChan2(ch1)
	for i := range ch2 {
		fmt.Println(i)
	}
}

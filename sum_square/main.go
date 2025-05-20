package main

import (
	"fmt"
	"sync"
)

func square(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i, v := range arr {
		arr[i] = v * v
	}
}
func sumOfArray(arr []int, wg *sync.WaitGroup) int {
	defer wg.Done()
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
func main() {
	var sum int
	var wg sync.WaitGroup
	arr := []int{2, 4, 6, 8, 10}
	wg.Add(1)
	go square(arr, &wg)
	wg.Wait()
	wg.Add(1)
	sum = sumOfArray(arr, &wg)
	wg.Wait()
	fmt.Println(sum)
}

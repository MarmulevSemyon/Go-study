package main

import (
	"fmt"
	"sync"
)

func main() {
	var sum int
	var wg sync.WaitGroup
	arr := []int{2, 4, 6, 8, 10}

	for _, v := range arr {
		wg.Add(1)
		go func(arr []int, v int) {
			defer wg.Done()
			sum += v * v
		}(arr, v)
	}

	wg.Wait()
	fmt.Println(sum)
}

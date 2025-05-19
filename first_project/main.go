package main

import (
	//"first_project/jopa"
	"fmt"
	//"github.com/go-sql-driver/mysql"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из слайс (2,4,6,8,10) и выведет их квадраты в stdout.
// func square(x int, wg *sync.WaitGroup) {

//		fmt.Println(x * x)
//		// почему так:
//		//wg.Done()
//		//а не  так (*wg).Done()
//	}
func main() {
	//var wg sync.WaitGroup
	arr := make([]int, 0)
	arr = append(arr, 1, 2, 3)
	fmt.Println(len(arr), cap(arr))

	// for _, v := range arr {
	// 	wg.Add(1)
	// 	go square(v, &wg)
	// }

	// wg.Wait()
}

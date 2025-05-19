package main

import (
	"fmt"
	"strconv"
)

func hj(n int, ch chan string) {
	defer close(ch)
	for i := 1; i <= n; i++ {

		if i%2 == 0 {
			ch <- strconv.Itoa(i) + " hui" // кладем данные в канал
		} else {
			ch <- strconv.Itoa(i) + " jopa" // кладем данные в канал
		}
	}
}

func myFunc(n int, ch chan int) {
	defer close(ch)
	for i := 1; i <= n; i++ {
		fmt.Println("положили данные в канал " + strconv.Itoa(i) + "-й раз")
		ch <- i
	}
	fmt.Println("канал закрыт (поидее), типо цикл закнчен дальше будет вызываться то, что в defer")
}
func main() {
	strCh := make(chan string)

	go hj(7, strCh)
	fmt.Println("first")
	for { //бесконечный цикл
		str, opened := <-strCh // получаем данные из канала
		if !opened {
			break // если канал закрыт, выходим из цикла
		}
		fmt.Println(str)
	}
	mych := make(chan int)
	go myFunc(7, mych)
	for {

		num, flag := <-mych
		if !flag {
			fmt.Println("канал закрыт в горутине main")
			break
		}
		fmt.Println("\t взяли данные из канала в горутине main :", num)
	}
}

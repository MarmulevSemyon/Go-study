package main

import (
	"fmt"
	"time"
)

func main() {
	ch_stop := make(chan int)

	go func_with_somthing_work(ch_stop) //запустили работу с каналом

	time.Sleep(4 * time.Second) // ждем какоете врямя
	close(ch_stop)
}
func func_with_somthing_work(ch chan int) {
	tick := time.NewTicker(time.Second)
	i := 1
	for {
		select {
		case <-tick.C:
			fmt.Printf("%d я нашел себе ррработку\nбуду чистить сковорррродку\n", i)
			i++
		case <-ch:
			return
		}
	}
}

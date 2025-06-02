package main

import (
	"fmt"
	"sync"
	"time"
)

// если в первая горутина положит в канал значение, и в это же время закроется канал ch_stop,
// то вторая горутина завершится и не будет забирать значение из канала, поэтому первая будет ждать,
// когда у неё заберут значение, а горутина main будет ждать завершения первой горутины (wg.Wait())
// можно убрать wg.Wait(), либо сделать буффер, тогда первая положит в буфер и сможет завершиться, либо как-то по-другому работать с select
func main() {
	ch_stop := make(chan int)
	ch := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go sentToChan(ch, ch_stop, &wg)
	go takeFromchan(ch, ch_stop, &wg)
	time.Sleep(1 * time.Second) // ждем какоето врямя
	close(ch_stop)
	wg.Wait()
}
func sentToChan(ch chan int, stop_ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	//defer close(ch)

	tick := time.NewTicker(3 * time.Millisecond)

	i := 1
	for {
		select {
		case <-tick.C:
			//fmt.Printf("%d я нашел себе ррработку\nбуду чистить сковорррродку\n", i)
			ch <- i
			i++
		case <-stop_ch:
			//tick.Stop()
			return
		}
	}
}
func takeFromchan(ch chan int, stop_ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("%d пришло\n", i)
		case <-stop_ch:
			return
		}
	}
}

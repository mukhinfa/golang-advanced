package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	chSlice := make(chan int)
	chSqrs := make(chan int)
	sl := make([]int, 0, 10)

	go func() {
		for range 10 {
			sl = append(sl, rand.Intn(100))
		}
		for _, v := range sl {
			chSlice <- v
		}
		close(chSlice)
	}()
	go func() {
		for v := range <-chSlice {
			chSqrs <- v * v
		}
		close(chSqrs)
	}()
	for {
		select {
		case p, ok := <-chSqrs:
			if !ok {
				return
			}
			fmt.Println(p)
		default:
			continue
		}
	}
}

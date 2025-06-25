package main

import (
	"fmt"
	"math/rand"
)

func main() {
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
		for range 10 {
			v := <-chSlice
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

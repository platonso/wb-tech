package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, num := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			square := num * num
			fmt.Println(square)
		}()
	}
	wg.Wait()
}

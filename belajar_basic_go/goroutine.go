package main

import (
	"fmt"
	"sync"
	"time"
)

func berhitung(a int, b int, wg *sync.WaitGroup, w int) {
	defer wg.Done()
	for a <= b {
		fmt.Println(a)
		a++
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("SELESAIIIIIIIII")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go berhitung(0, 100, &wg, 100)
	go berhitung(1, 100, &wg, 400)
	go berhitung(2, 100, &wg, 400)

	wg.Wait()
	fmt.Println("semua selesaiiiiiii")
}

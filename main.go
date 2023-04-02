package main

import (
	"go_design_pattern/design_pattern"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go design_pattern.GetInstance4(&wg)
	}
	wg.Wait()
}

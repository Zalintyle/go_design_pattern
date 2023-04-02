package main

import (
	"design_pattern_learning/design_pattern"
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

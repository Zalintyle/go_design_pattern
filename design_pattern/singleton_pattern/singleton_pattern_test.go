package singleton_pattern

import (
	"sync"
	"testing"
)

func TestGetInstance4(t *testing.T) {
	// test singleton pattern
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go GetInstance4ForTest(&wg)
	}
	wg.Wait()
}

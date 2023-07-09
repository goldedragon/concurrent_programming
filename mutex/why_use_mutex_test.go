package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestCount(t *testing.T) {
	// 使用mutex解决count++结果的问题
	var (
		count = 0
		wg    sync.WaitGroup
		lock  sync.Mutex
	)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				lock.Lock()
				count++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

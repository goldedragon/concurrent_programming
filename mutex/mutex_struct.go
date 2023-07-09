package main

import (
	"fmt"
	"sync"
)

// 可以在结构体中使用
type counter struct {
	count int
	mu    sync.Mutex
}

// 还可以把mutex封装成一个方法
func (c *counter) add() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func main() {
	var (
		c  counter
		wg sync.WaitGroup
	)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				/*c.mu.Lock()
				c.count++
				c.mu.Unlock()
				*/
				c.add()
			}
		}()
	}
	wg.Wait()
	fmt.Println(c.count)
}

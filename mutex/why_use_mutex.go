package main

import (
	"fmt"
	"sync"
)

// 为什么一定要加锁
func main() {
	/*
		我们创建了 10 个 goroutine，同时不断地对一个变量（count）进行加 1 操作
		每个 goroutine 负责执行 10 万次的加 1 操作，
		我们期望的最后计数的结果是 10 * 100000 = 1000000 (一百万)。
	*/
	var count = 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				// count++不是原子操作
				count++ // 临界区
				// 分为三步
				// 1.读取count当前值
				// 2.对读取到的count值+1
				// 3.把结果写回到count中
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
	// 结果：差的太多

}

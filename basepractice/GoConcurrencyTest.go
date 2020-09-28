package main

import (
	"fmt"
	"time"
)

func Producer(begin, end int, queue chan <- int)  {
	for i := begin ; i < end; i++ {
		fmt.Println("producer ", i)
		queue <- i
	}
}

func Consumer(queue <- chan  int)  {
	for val := range queue {
		fmt.Println("consumer ", val)
	}
}

func main() {
	queue := make(chan int)
	defer close(queue)
	for i := 0; i < 3; i++ {
		// 多个生产者
		go Producer(1 * 5, (i + 1) *5, queue)
	}
	go Consumer(queue) // 一个消费者
	time.Sleep(time.Second)

	/**
	注意 main函数执行结束，整个GO也会结束 无论是否有未执行完的goruntine
	 */
}

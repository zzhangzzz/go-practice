package main

import (
	"fmt"
	"time"
)

func send(ch chan int, begin int) {
	for i := begin ; i< begin + 10; i++ {
		ch <- i
	}
}

func receive(ch <-chan int) {
	val := <- ch
	fmt.Println("received :", val)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1, 0)
	go receive(ch2)

	time.Sleep(time.Second)

	/**
	由于 ch2 中的消息仅被接收一次，所以仅出现一次“send value by ch2”，后续消息的发送将被阻塞。
	select 语句分别从 3 个 case 中选取返回的 case 进行处理，
	当有多个 case 语句同时返回时，
	select 将会随机选择一个 case 进行处理。
	如果 select 语句的最后包含 default 语句，
	该 select 语句将会变为非阻塞型，
	即当其他所有的 case 语句都被阻塞无法返回时，
	select 语句将直接执行 default 语句返回结果。
	在最后的 case 语句使用了 <-time.After(2 * time.Second) 的方式指定了定时返回的 channel，
	这是一种有效从阻塞的 channel 中超时返回的小技巧。
	 */
	for {
		select {
		case val := <-ch1:
			fmt.Println("get value %d from ch1 \n", val)
		case ch2 <- 2:
			fmt.Println("send value by ch2")
		case <- time.After(2 * time.Second):
			fmt.Println("time out")
			return
		}
	}
}

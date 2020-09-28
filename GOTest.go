package main

import (
	"fmt"
	"sync"
)

func input(ch chan string) {
	defer wg.Done()
	defer close(ch)

	var input string
	fmt.Println("Enter EOF to shut down:")

	for {
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println("read input err:", err.Error())
			break
		}

		if input == "EOF" {
			fmt.Println("end!!")
			break
		}

		ch <- input
	}
}

func output(ch chan string)  {
	defer wg.Done()
	for value := range ch {
		fmt.Println("your input is :", value)
	}
}

var wg sync.WaitGroup

func main()  {
	ch := make(chan string)
	wg.Add(2)
	// 读取到输入
	go input(ch)
	// 输出到命令
	go output(ch)
	wg.Wait()
	fmt.Println("exit")
}
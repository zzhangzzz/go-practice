package main

import (
	"fmt"
	"time"
)

func add(a, b int) int {
	return a+b
}

func main() {
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}

	sum2 := 1
	for sum2 < 100 {
		sum2 *= 2
	}

	sum3 := 1
	for {
		if sum3 < 100 {
			sum3 *= 2
		} else {
			break
		}
	}
	fmt.Println(sum1, sum2, sum3)


	/**
	go 的switch不需要break fallthrough可以链接两个case的情况
	 */
	nowTime := time.Now()
	switch nowTime.Weekday() {
	case time.Saturday:
		fmt.Println("have a rest")
	case time.Sunday:
		fmt.Println("hava a rest too")
	default:
		fmt.Println("go to work")
	}

	switch {
	case nowTime.Weekday() >= time.Monday && nowTime.Weekday() <= time.Friday :
		fmt.Println("go to work")
	default:
		fmt.Println("have a rest")
	}

	/**
	延迟执行
	defer关键字会在return 返回前执行 按照先进后出的顺序调用
	可以用来进行资源释放等工作
	 */
	a := 1
	b := 2
	defer fmt.Println("result ", add(a, b))
	a = 3
	b = 4
	defer fmt.Println("result last ", add(a, b))

	a = 5
	b = 6
	fmt.Println(add(a, b))
}

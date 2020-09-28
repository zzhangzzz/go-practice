package main

import (
	"context"
	"fmt"
	"time"
)

const DB_ADDR = "db_addr"
const CALCULATE_VAL = "calculate_val"

func ReadDB(ctx context.Context, cost time.Duration)  {
	fmt.Println("db_addr is", ctx.Value(DB_ADDR))
	select {
	case <- time.After(cost):
		fmt.Println("read data from db")
	case <- ctx.Done():
		fmt.Println(ctx.Err())

	}
}

func calculate(ctx context.Context, cost time.Duration)  {
	fmt.Println("calculate value is ", ctx.Value(CALCULATE_VAL))
	select {
	case <- time.After(cost) :
		fmt.Println("calculate finish")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {
	// 创建上下文
	ctx := context.Background()
	// 添加上下文
	ctx = context.WithValue(ctx,  DB_ADDR, "localhost:1222")
	ctx = context.WithValue(ctx, CALCULATE_VAL, 1234)

	ctx,cancel := context.WithTimeout(ctx, time.Second * 2)
	defer cancel()
	go ReadDB(ctx, time.Second * 4)
	go calculate(ctx, time.Second * 4)

	time.Sleep(time.Second * 5)
}

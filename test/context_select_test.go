package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextSelectTest(t *testing.T) {
	// 创建一个带有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	ch := make(chan int)

	// 模拟一个耗时的操作，通过通道传递结果
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 42
	}()

	// 使用 select 监听通道和 context
	select {
	case result := <-ch:
		fmt.Println("接收到结果:", result)
	case <-ctx.Done():
		fmt.Println("超时或取消")
	}

	fmt.Println("执行到最后select了")
}

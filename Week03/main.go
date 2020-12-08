package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eg, ctx := errgroup.WithContext(context.Background())
	// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
	eg.Go(func() error {
		fmt.Println("goroutine1: service start")
		select {
			case <- ctx.Done():
				fmt.Println("goroutine1: service start goroutine cancel")
				return ctx.Err()
		}
		return nil
	})

	eg.Go(func() error {
		fmt.Println("goroutine2: service stop")
		for{
			select {
				case <- ctx.Done():
					fmt.Println("goroutine2: service stop goroutine cancel")
					return ctx.Err()
			}
		}
		return nil
	})

	eg.Go(func() error {
		fmt.Println("goroutine3: signal")
		sings := make(chan os.Signal)
		signal.Notify(sings, syscall.SIGINT)
		for{
			select {
				case <-ctx.Done():
					fmt.Println("goroutine3: signal goroutine cancel")
					return ctx.Err()
				case <-sings:
					fmt.Println("goroutine3: signal goroutine cancel")
					return errors.New("goroutine3: signal hand exit")
			}
		}
		return nil
	})

	err := eg.Wait()
	if err != nil {
		fmt.Printf("err: %+v", err)
		return
	}
	fmt.Println("done")
}


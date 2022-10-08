package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(6*time.Second))
	defer cancel()

	go func2(ctx, cancel)
	go func() {
		err := func1(ctx, cancel)
		fmt.Println("action 1 err:", err)
		if err != nil {
			println(err.Error())
		}
	}()

	go func() {
		i := 0

		for {
			i++
			time.Sleep(1 * time.Second)
			fmt.Println(i)
			select {
			case <-ctx.Done():
				fmt.Println("ctx.Done")
				os.Exit(1)
			default:
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)
	<-quit

	fmt.Println("final done.")
}

func func1(ctx context.Context, cancelFun context.CancelFunc) error {
	fmt.Println("action 1")

	time.Sleep(5 * time.Second)
	return errors.New("failed")
}

func func2(ctx context.Context, cancelFun context.CancelFunc) {
	fmt.Println("action 2")
	time.Sleep(7 * time.Second)
	cancelFun()
	fmt.Println("action 2 done")
}

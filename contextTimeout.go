package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	i := 0
	go func() {
		for {
			fmt.Println(i)
			i++

			time.Sleep(1 * time.Second)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// ctx, cancel := context.WithDeadline(context.Background(), <-time.After(5*time.Second))
	defer cancel()

d:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done.")
			break d
		default:
			fmt.Println("in progress...")
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("Stop.")
}

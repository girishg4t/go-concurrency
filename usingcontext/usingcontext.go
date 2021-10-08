package usingcontext

import (
	"context"
	"fmt"
	"time"
)

func StopExecution() {
	ctx, cancel := context.WithCancel(context.Background())
	go one(ctx)
	go two(ctx)
	go three(ctx)
	go four(ctx)
	select {
	case <-time.After(time.Second * 10):
		fmt.Println("All done.")
		cancel()
	}
}

func one(ctx context.Context) {
	go func() {
		var isDone bool
		for {
			if !isDone {
				fmt.Println("one")
			}
			<-ctx.Done()
		}
	}()
}

func four(ctx context.Context) {
	go func() {
		var isDone bool
		for {
			if !isDone {
				fmt.Println("four")
			}
			<-ctx.Done()
		}
	}()
}
func two(ctx context.Context) {
	go func() {
		var isDone bool
		for {
			if !isDone {
				fmt.Println("two")
			}
			<-ctx.Done()
		}
	}()
}
func three(ctx context.Context) {
	go func() {
		var isDone bool
		for {
			if !isDone {
				fmt.Println("three")
			}
			<-ctx.Done()
		}
	}()
}

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
		for {
			select {
			case <-ctx.Done():
				break
			default:
				fmt.Println("Started One")
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	}()
}

func four(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				fmt.Println("Started Four")
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	}()
}
func two(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				fmt.Println("Started Two")
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	}()
}
func three(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				fmt.Println("Started Three")
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	}()
}

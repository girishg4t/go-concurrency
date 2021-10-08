package selectchannels

import (
	"fmt"
	"time"
)

func StopExecution() {

	done := make(chan bool)
	go one(done)
	go two(done)
	go three(done)
	go four(done)
	select {
	case <-time.After(time.Second * 10):
		fmt.Println("All done.")
		done <- true
	}
}

func one(done chan bool) {
	go func() {
		var isDone bool
		for {
			if !isDone {
				fmt.Println("Started One")
			}
			isDone = <-done
		}
	}()
}

func four(done chan bool) {
	go func() {
		var isDone bool
		for {
			if !isDone {
				fmt.Println("Started Four")
			}
			isDone = <-done
		}
	}()
}
func two(done chan bool) {
	go func() {
		var isDone bool
		for {
			if !isDone {
				fmt.Println("Started Two")
			}
			isDone = <-done
		}
	}()
}
func three(done chan bool) {
	go func() {
		var isDone bool
		for {
			if !isDone {
				fmt.Println("Started Three")
			}
			isDone = <-done
		}
	}()
}

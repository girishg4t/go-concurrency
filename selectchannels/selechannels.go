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
		for {
			select {
			case <-done:
				break
			default:
				fmt.Println("Started One")
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	}()
}

func four(done chan bool) {
	go func() {
		for {
			select {
			case <-done:
				break
			default:
				fmt.Println("Started Four")
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	}()
}
func two(done chan bool) {
	go func() {
		for {
			select {
			case <-done:
				break
			default:
				fmt.Println("Started Two")
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	}()
}
func three(done chan bool) {
	go func() {
		for {
			select {
			case <-done:
				break
			default:
				fmt.Println("Started Three")
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	}()
}

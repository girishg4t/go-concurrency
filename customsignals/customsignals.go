package customsignals

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func timeBound(n int, done chan bool) {
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for x := 0; x < n; x++ {
		fmt.Fprintf(f, ".")
	}
	select {
	case <-time.After(time.Second * 10):
		done <- true
	}
}

func Start() {
	start := time.Now()
	var pool = 2
	const numJobs = 100
	currentjob := 1
	//var signal chan int = make(chan int)
	signals := make(chan os.Signal, 1)
	for currentjob < numJobs {
		signal.Notify(signals, os.Interrupt)
		go func(size, job int) {
			fmt.Printf("job %d\n", job)
			newPool(size, job)
		}(pool, currentjob)
		currentjob = currentjob + pool
		<-signals
		fmt.Println("signal received")

		if pool == 2 {
			pool = 1
		} else {
			pool = 2
		}
	}

	duration := time.Since(start).Seconds()
	fmt.Print(fmt.Sprintf("%v", duration))
}

func newPool(pool int, job int) {
	fmt.Printf("Pool increased/decreased by 1. Current Pool size: %d\n", pool)

	messages := make(chan bool, pool)
	for w := 1; w <= pool; w++ {
		timeBound(w, messages)
	}
	<-messages

}

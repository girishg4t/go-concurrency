package pool

import (
	"fmt"
	"os"
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

	const numJobs = 10
	job := 1

	for job < numJobs {
		messages := make(chan bool, 2)
		for w := 1; w <= 2; w++ {
			go timeBound(w, messages)
		}
		job = job + 2
		<-messages
	}

	duration := time.Since(start).Seconds()
	fmt.Print(fmt.Sprintf("%v", duration))
}

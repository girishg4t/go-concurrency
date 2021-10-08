package selectmultychannels

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/girishg4t/go-concurrency/utils"
)

func Start() {
	nums := 8
	for i := 0; i < nums; i++ {
		go utils.CpuBound(i + 1)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
		fmt.Println("Caught interrupt")
	case <-time.After(10 * time.Second):
		fmt.Println("Goodbye! timeout")
	}
	fmt.Println("All done ")
}

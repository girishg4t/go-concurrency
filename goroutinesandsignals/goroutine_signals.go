package goroutinesandsignals

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/girishg4t/go-concurrency/utils"
)

func cpuConsumed(n, rn int) {
	go utils.CpuBound(rn)
}

func Start() {
	idle0, total0 := utils.GetCPUSample()
	go utils.CpuBound(0)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
	fmt.Println("Goodbye")
	idle1, total1 := utils.GetCPUSample()
	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	fmt.Printf("CPU consumed by this process %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)

}

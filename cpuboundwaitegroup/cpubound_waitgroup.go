package cpubountwaitegroup

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"

	"github.com/girishg4t/go-concurrency/utils"
)

func cpuConsumed(n, rn int) {
	go utils.CpuBound(rn)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
	fmt.Println(fmt.Sprintf("Goodbye %d", n))
}

func Start() {
	idle0, total0 := utils.GetCPUSample()
	nums := 16
	for i := 0; i < nums; i++ {
		ix := rand.Intn(i + 1)
		cpuConsumed(i, ix)
	}
	idle1, total1 := utils.GetCPUSample()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)

}

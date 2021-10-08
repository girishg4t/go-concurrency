package fanin

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/girishg4t/go-concurrency/utils"
)

func (f FanInService) StartWithoutMutex() {
	start := time.Now()
	var arr []int = make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = getRandomNumber(0, 10)
	}
	var numMap map[int]int = make(map[int]int)
	for _, n := range arr {
		output := utils.MathBound(n)
		numMap[n] = output
	}
	duration := time.Since(start).Seconds()
	fmt.Println(fmt.Sprintf("%v", duration))
	fmt.Println(numMap)
}

func getRandomNumber(min, max int) int {
	return (rand.Intn(max-min) + min)
}

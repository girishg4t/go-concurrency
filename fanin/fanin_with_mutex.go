package fanin

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type FanInService struct {
	app *sync.Mutex
}

func NewFanInService() FanInService {
	return FanInService{
		app: &sync.Mutex{},
	}
}

func (f FanInService) mathBoundWithMutex(n int) int {
	time.Sleep(time.Duration(n) * time.Second)
	result := n * 8
	return result
}

func (f FanInService) StartWithMutex() {
	start := time.Now()
	var arr []int = make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = getRandomNumber(0, 10)
	}
	var numMap map[int]int = make(map[int]int)
	for _, n := range arr {
		f.app.Lock()
		fmt.Println(n)
		numMap[n] = f.mathBoundWithMutex(n)
		f.app.Unlock()
	}
	duration := time.Since(start).Seconds()
	fmt.Println(fmt.Sprintf("%v", duration))
	fmt.Println(numMap)
}

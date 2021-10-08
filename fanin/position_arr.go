package fanin

import (
	"fmt"
	"time"
)

func (f FanInService) StartWithMutexOutArr() {
	start := time.Now()
	var arr []int = make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = getRandomNumber(0, 10)
	}
	var result []int = make([]int, 0)
	var numMap map[int]int = make(map[int]int)
	for _, n := range arr {
		f.app.Lock()
		_, exist := numMap[n]
		if exist {
			result = append(result, n)
		}
		numMap[n] = f.mathBoundWithMutex(n)
		f.app.Unlock()
	}
	duration := time.Since(start).Seconds()
	fmt.Println(fmt.Sprintf("%v", duration))
	fmt.Println(result)
}

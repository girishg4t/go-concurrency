package main

import (
	"flag"

	cb "github.com/girishg4t/go-concurrency/cpuboundwaitegroup"
	"github.com/girishg4t/go-concurrency/customsignals"
	"github.com/girishg4t/go-concurrency/fanin"
	sig "github.com/girishg4t/go-concurrency/goroutinesandsignals"
	"github.com/girishg4t/go-concurrency/pool"
	"github.com/girishg4t/go-concurrency/selectchannels"
	"github.com/girishg4t/go-concurrency/selectmultychannels"
	"github.com/girishg4t/go-concurrency/usingcontext"
)

func main() {
	program := flag.Int("program", 0, "which program to start")
	innerProg := flag.Int("inner-program", 0, "which inner program to start")
	flag.Parse()
	switch *program {
	case 1:
		sig.Start()
	case 2:
		cb.Start()
	case 3:
		selectchannels.StopExecution()
	case 4:
		usingcontext.StopExecution()
	case 5:
		pool.Start()
	case 6:
		selectmultychannels.Start()
	case 7:
		fis := fanin.NewFanInService()
		if *innerProg == 1 {
			fis.StartWithMutex()
		}
		if *innerProg == 2 {
			fis.StartWithoutMutex()
		}
		if *innerProg == 3 {
			fis.StartWithMutexOutArr()
		}

	case 8:
		customsignals.Start()
	}

}

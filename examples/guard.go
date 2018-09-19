package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ninedraft/gocontrol"
)

func main() {
	var guard = &gocontrol.Guard{}
	var wg = &sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		i := i

		wg.Add(1)
		go func() {
			defer guard.Go()()
			defer wg.Done()
			fmt.Printf("START %d\n", i)
			time.Sleep(100 * time.Duration(i) * time.Millisecond)
			fmt.Printf("END %d\n", i)
		}()
	}

	var done = doAndClose(wg.Wait)
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("%d goroutines are running\n", guard.AliveN())
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func doAndClose(do func()) <-chan struct{} {
	var done = make(chan struct{})
	go func() {
		do()
		close(done)
	}()
	return done
}

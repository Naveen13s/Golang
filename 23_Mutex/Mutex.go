//A Mutex (Mutual Exclusion) is used to protect shared data from being accessed by multiple goroutines at the same time.
//Without a mutex, concurrent reads/writes can cause race conditions.

package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int // the movement it is modified we will lock this, that means once in a time goroutine will able to modified only one, by using Mutex.
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1

}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	//
	//myPost.inc()
	//myPost.inc()
	fmt.Println(myPost.views)

}

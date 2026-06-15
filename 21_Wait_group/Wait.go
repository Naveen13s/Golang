// A WaitGroup from Go's sync package is used to wait for a collection of goroutines to finish executing.
// It helps ensure that the main function doesn't exit before all goroutines complete.



package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
		/*go func(i int) {
			fmt.Println(i)
		}(i) */
	}
	wg.Wait()

}

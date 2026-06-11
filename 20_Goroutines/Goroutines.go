//A goroutine is a lightweight thread managed by the Go runtime.
//It allows us to run functions concurrently using the go keyword.
// Without time.Sleep(), the program might exit before the goroutine runs.

/* Goroutines:
  ->  Start with a very small stack (~2 KB)
  ->  Managed by Go runtime
->    Thousands of goroutines can run efficiently   */

package main

import (
	"fmt"
)

//func task(id int) {
//	fmt.Println("doing task", id)
//}
/*
func main() {
	for i := 0; i <= 10; i++ {
		//go task(i)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 2)

}  */

// Channels
func calculate(ch chan int) {
	ch <- 100
}

func main() {
	ch := make(chan int)
	go calculate(ch)
	result := <-ch
	fmt.Println(result)
}

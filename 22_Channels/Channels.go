/* A channel is a way for goroutines to communicate and share data safely.
   ->"Don't communicate by sharing memory; share memory by communicating." — Go Philosophy
   -> Channels enable goroutine communication
   -> Unbuffered channels are blocking
   -> Buffered channels store values temporarily
   -> close() signals no more values will be sent
   -> select works with multiple channels
   -> Channels are Go's preferred concurrency mechanism

*/

package main

import "fmt"

/*
func main() {
	messageChan := make(chan string) // channel made

	messageChan <- "ping"  // data send

	msg := <-messageChan // data receive

	fmt.Println(msg)

}  /* output  -  PS D:\Golang> go run 22_Channels/Channels.go
       fatal error: all goroutines are asleep - deadlock!

       goroutine 1 [chan send]:
       main.main()
        D:/Golang/22_Channels/Channels.go:10 +0x36
        exit status 2

*/
/* // Send

func processNum(numChan chan int) {

	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)

	}




}
func main() {
	numChan := make(chan int)

	go processNum(numChan)

	for {
		numChan <- rand.Intn(100)
	}

	//numChan <- 5

	//time.Sleep(time.Second * 2)
}

//func main() {
//messageChan := make(chan string)

//messageChan <- "ping" // blocking

//msg := <-messageChan //

//fmt.Println(msg)

//}  */

/*  To receive Data
func sum(result chan int, num1 int, num2 int) {

	numResult := num1 + num2
	result <- numResult

}

func main() {

	result := make(chan int)

	go sum(result, 4, 5)

	res := <-result  // blocking
	fmt.Println(res)

}  */

/* the work whcih has done by using waitgroup we can also do by using channels.
// goroutine synchronizer
func task(done chan bool) {
	defer func() { done <- true }()

	fmt.Println("processing....")
	//done <- true  //last line

}

func main() {

	done := make(chan bool)
	go task(done)

	<-done // block

}  */
/*
// for synchronization when we need to use channel or waitgroup then for single go routine we can go with channel and for multiple  wait group is best.
// q
//Single Chnnel:
func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sneding email to ", email)
		time.Sleep(time.Second)
	}

}
func main() {

	emailChan := make(chan string, 10)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending.")
	close(emailChan) // closing the channnel is very imp otherwiswe deadlock will achieve
	<-done

}  */

/*
// buffered channel

func main() {

	emailChan := make(chan string, 100)

	emailChan <- "1@example.com"
	emailChan <- "2@example.com"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)

}  */

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {

		select {
		case chan1Val := <-chan1:
			fmt.Println("receive data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("receive data from chan2", chan2Val)
		}
	}
}

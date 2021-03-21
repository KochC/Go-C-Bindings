package main

import (
	"fmt"
	"time"
)

// #cgo CFLAGS: -Wall -I${SRCDIR}/
// #cgo LDFLAGS: -L${SRCDIR}/ -ltest_c
// #include "test.h"
import "C"

func main() {

	go c_code_handler()
	time.Sleep(time.Millisecond * 500)
	go c_code_handler()
	fmt.Scanln()

	// press any key to proceed to worker example

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++{
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++{
		fmt.Println(<-results)
	}
}

func c_code_handler(){
	number := C.int(0)
	for{
		result := C.increment((C.int)(number))
		number = result
		fmt.Println(number)
		time.Sleep(time.Second * 1)
	}
}

// only receiving on jobs channel, only send on result channel
func worker(jobs <-chan int, result chan<- int){
	for n := range jobs{
		result <- fib(n)
	}
}

func fib(n int) int{
	if n <= 1{
		return n
	}
	return fib(n-1) + fib(n-2)
}
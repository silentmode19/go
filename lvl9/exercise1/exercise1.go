/* In addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists */

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Hi there")
	fmt.Println(runtime.NumCPU())
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("Hello from subroutine")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from subroutine2")
		wg.Done()
	}()
	wg.Wait()
}

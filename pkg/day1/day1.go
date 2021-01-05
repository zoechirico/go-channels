package day1

import (
	 "sync"
	 "fmt"
	 "time"
)


func UsingWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("This is first")
		time.Sleep(1)
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("This is 2nd")
		time.Sleep(1)
	}()


	wg.Wait()
}
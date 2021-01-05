package day1

import (
	"testing"
	"sync"
	"fmt"
	"time"
	//"fmt"
)


func Test_UsingWaitGroup(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()
		fmt.Println("We want something to say 1")
		time.Sleep(1)
	}()


	wg.Add(1)

	go func() {

		defer wg.Done()
		fmt.Println("We want something to say 2")
		time.Sleep(1)
	}()

	wg.Wait()

}

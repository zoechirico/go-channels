package syncpools

import (
	"testing"
	"sync"
)

func Test_thing(t *testing.T) {

	var wg  sync.WaitGroup

	const N = 3000
	wg.Add(2*N)
    for i:=0; i < N; i++ {
		go func() {
			defer wg.Done()
			Example1()
		}()
		go func() {
			defer wg.Done()
			Example2()
		}()
	 
	}
	wg.Wait()

}

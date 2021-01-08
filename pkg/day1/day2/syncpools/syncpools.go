package syncpools

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

var pool = sync.Pool{
	// New creates an object when the pool has nothing available to return.
	// New must return an interface{} to make it flexible. You have to cast
	// your type after getting it.
	New: func() interface{} {
		// Pools often contain things like *bytes.Buffer, which are
		// temporary and re-usable.
		return &bytes.Buffer{}
	},
}

func Example1() {

	// When getting from a Pool, you need to cast
	s := pool.Get().(*bytes.Buffer)
	// We write to the object
	s.Write([]byte("EXAMPLE 1 "))

	// Then put it back
	pool.Put(s)

	time.Sleep(3 * time.Millisecond)
	// Pools can return dirty results

	// Get 'another' buffer
	s = pool.Get().(*bytes.Buffer)
	// Write to it
	s.Write([]byte("EXAMPLE 1 (APPEND) "))
	// At this point, if GC ran, this buffer *might* exist already, in
	// which case it will contain the bytes of the string "dirtyappend"
	fmt.Println(s)
	// So use pools wisely, and clean up after yourself
	s.Reset()
	pool.Put(s)

	// When you clean up, your buffer should be empty
	s = pool.Get().(*bytes.Buffer)
	// Defer your Puts to make sure you don't leak!
	defer pool.Put(s)
	s.Write([]byte("EXAMPLE1 (RESET)"))
	// This prints "reset!", and not "dirtyappendreset!"
	fmt.Println(s)
}

func Example2() {

	// When getting from a Pool, you need to cast
	s := pool.Get().(*bytes.Buffer)
	// We write to the object
	s.Write([]byte("Exampe2 ->A <-"))
	// Then put it back
	pool.Put(s)

	// Pools can return dirty results

	// Get 'another' buffer
	s = pool.Get().(*bytes.Buffer)
	// Write to it
	s.Write([]byte("Example2 ->A append<-"))
	// At this point, if GC ran, this buffer *might* exist already, in
	// which case it will contain the bytes of the string "dirtyappend"
	fmt.Println(s)
	// So use pools wisely, and clean up after yourself
	s.Reset()
	pool.Put(s)

	// When you clean up, your buffer should be empty
	s = pool.Get().(*bytes.Buffer)
	// Defer your Puts to make sure you don't leak!
	defer pool.Put(s)
	s.Write([]byte("Example2 (reset!)"))
	time.Sleep(3 * time.Microsecond)
	s.Reset()
	// This prints "reset!", and not "dirtyappendreset!"
	fmt.Println(s)
}

func Z() {

	fmt.Printf("zoe")

	s := pool.Get().(*bytes.Buffer)
	s.Write([]byte("func Z"))
	defer s.Put()

}

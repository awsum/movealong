// run with go test -test.bench Bench

package main

import (
	// "fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

const nProcuder = 4 // better to set this to core number
const nItem = 2000000
const totalItem = nProcuder * nItem

var data = [totalItem]byte{}

func init() {
	runtime.GOMAXPROCS(nProcuder)
}

// channel version
//var channel = make(chan byte, nProcuder)

var channel = make(chan byte, totalItem)

func produce() {
	for i := 0; i < nItem; i++ {
		channel <- 1
	}
}

func consume() {
	for i := 0; i < nItem; i++ {
		data[i] = <-channel
	}
}

func testChannel() {
	//fmt.Println("start test chan,", nProcuder, "producer")
	for i := 0; i < nProcuder; i++ {
		go produce()
	}
	consume()
}

func BenchmarkChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testChannel()
	}
}

// lock version

var mutex sync.Mutex
var id int

var wait int32

func addData(wg *sync.WaitGroup) {
	atomic.AddInt32(&wait, 1)
	for wait != nProcuder {
		// wait all go routine start doing actual work
	}
	for i := 0; i < nItem; i++ {
		mutex.Lock()
		data[id] = 1
		id++
		mutex.Unlock()
	}
	wg.Done()
}

func testLock() {
	wait = 0
	id = 0
	var wg sync.WaitGroup
	wg.Add(nProcuder)
	//fmt.Println("start test lock,", nProcuder, "producer")
	for i := 0; i < nProcuder; i++ {
		go addData(&wg)
	}
	wg.Wait()
}

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testLock()
	}
}

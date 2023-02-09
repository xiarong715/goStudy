package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter interface {
	Inc()
	Load() int64
}

type CommonCounter struct {
	counter int64
}

func (c *CommonCounter) Inc() {
	c.counter++
}

func (c *CommonCounter) Load() int64 {
	return c.counter
}

type MutextCounter struct {
	counter int64
	mutext  sync.Mutex
}

func (mc *MutextCounter) Inc() {
	mc.mutext.Lock()
	defer mc.mutext.Unlock()
	mc.counter++
}

func (mc *MutextCounter) Load() int64 {
	mc.mutext.Lock()
	defer mc.mutext.Unlock()
	return mc.counter
}

type AtomicCounter struct {
	counter int64
}

func (ac *AtomicCounter) Inc() {
	atomic.AddInt64(&ac.counter, 1)
}

func (ac *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&ac.counter)
}

func test(c Counter) {
	wg := sync.WaitGroup{}
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("c.counter = %v, time = %v\n", c.Load(), end.Sub(start))
}

package stoppable

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestStoppable_NewStopable(t *testing.T) {
	now := time.Now()
	n := 100000
	s := NewStopable()
	for i := 0; i < n; i++ {
		if !s.Alive() {
			t.Errorf("Should not be dead [%d]", i)
		}
	}
	for i := 0; i < n; i++ {
		s.Die()
	}

	for i := 0; i < n; i++ {
		if s.Alive() {
			t.Errorf("Not dead [%d]", i)
		}
	}
	tt := time.Since(now)
	fmt.Printf("Took %v, or %v per op\n", tt, tt/time.Duration(n))
}

func TestStoppable_Many(t *testing.T) {
	n := 100000
	stopat := n >> 1
	s := NewStopable()
	wg := sync.WaitGroup{}

	start := func(i int) {
		if i >= stopat {
			s.Die()
		}
		s.Alive()
		wg.Done()
	}

	now := time.Now()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go start(i)
	}
	wg.Wait()
	tt := time.Since(now)
	fmt.Printf("Took %v, or %v per op\n", tt, tt/time.Duration(n))
}

package stoppable

import (
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
	t.Logf("Took %v, or %v per op", tt, tt/time.Duration(n))
}

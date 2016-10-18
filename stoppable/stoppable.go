/*
MIT License

Copyright (c) 2016 Nick Potts

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package stoppable

import (
	"sync"
)

/*Halter is a design pattern in that something can be started,
and permanently stopped in an asynchronous and race-free way.*/
type Halter interface {
	Alive() bool //returns true if alive, false otherwise
	Die()        //kills the process
}

/*Stoppable is a stoppable design pattern in that something can be started,
and permanently stopped in an asynchronous and race-free way.*/
type Stopable struct {
	alive chan bool
	kill  chan error
	m     sync.Mutex
}

/*monitor emits alive bits until the channel is closed. It will close all channels
before it exists*/
func (s *Stopable) monitor() {
	<-s.kill //sync with new()
	for {
		select {
		case <-s.kill:
			s.kill <- nil //lock
			close(s.alive)
			close(s.kill)
			return
		case s.alive <- true:
		}
	}
}

/*Alive returns true if it is still alive,  otherwise dead*/
func (s *Stopable) Alive() bool {
	return <-s.alive
}

/*Die marks the the process as dead*/
func (s *Stopable) Die() {
	s.m.Lock()
	defer s.m.Unlock()
	if s.Alive() {
		s.kill <- nil
		<-s.kill //lockstep with monitor
		return
	}
}

/*NewStopable creates an instance of a stoppable process*/
func NewStopable() Halter {
	r := &Stopable{
		alive: make(chan bool),
		kill:  make(chan error),
	}
	go r.monitor()
	r.kill <- nil
	return r
}

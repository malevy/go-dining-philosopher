package main

import "time"

const LockToken int = 1

type Fork struct {
	label string
	lock  chan int
}

func NewFork(label string) *Fork {
	return &Fork{label: label, lock: make(chan int, 1)}
}

func (f Fork) TryPickup(d time.Duration) bool {

	select {
	case f.lock <- LockToken:
		return true
	case <-time.After(d):
		return false
	}
}

func (f Fork) PutDown() {
	<-f.lock
}

package main

import (
	"fmt"
	"strconv"
)

func main() {

	const philosopherCount = 5

	forks := make([]*Fork, philosopherCount)
	for i := 0; i < len(forks); i++ {
		forks[i] = NewFork(fmt.Sprintf(strconv.Itoa(i)))
	}

	philosophers := make([]*Philosopher, philosopherCount)
	for i := 0; i < len(philosophers); i++ {
		philosophers[i] = NewPhilosopher(strconv.Itoa(i), forks[i], forks[(i+1)%philosopherCount])
	}

	for _, p := range philosophers {
		go p.Dine()
	}

	fmt.Scanln()

}

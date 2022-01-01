package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Philosopher struct {
	label               string
	leftFork, rightFork *Fork
}

func (p Philosopher) String() string {
	return p.label
}

func NewPhilosopher(label string, leftFork *Fork, rightFork *Fork) *Philosopher {
	return &Philosopher{label: label, leftFork: leftFork, rightFork: rightFork}
}

func (p Philosopher) Dine() {

	var (
		randomSource              = rand.New(rand.NewSource(time.Now().UnixNano()))
		activityDurationGenerator = durationGeneratorGenerator(randomSource, 100)
		lockWaitDurationGenerator = durationGeneratorGenerator(randomSource, 10)
	)

	for true {
		fmt.Printf("philosopher %v is thinking\n", p)
		time.Sleep(activityDurationGenerator())

		if !p.leftFork.TryPickup(lockWaitDurationGenerator()) {
			fmt.Printf("philosopher %v failed to acquire left fork\n", p)
			continue
		}
		if !p.rightFork.TryPickup(lockWaitDurationGenerator()) {
			fmt.Printf("philosopher %v failed to acquire right fork\n", p)
			p.leftFork.PutDown()
			continue
		}

		fmt.Printf("philosopher %v is eating\n", p)
		time.Sleep(activityDurationGenerator())

		p.rightFork.PutDown()
		p.leftFork.PutDown()

	}

}

func durationGeneratorGenerator(r *rand.Rand, factor int) func() time.Duration {
	return func() time.Duration {
		return time.Duration(r.Intn(factor)) * time.Millisecond
	}
}

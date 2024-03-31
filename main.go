package lab1

import (
	"fmt"
	"sync"
	"time"
)

var eatGroup sync.WaitGroup

type fork struct{ sync.Mutex }

type Philosopher struct {
	id                  int
	leftFork, rightFork *fork
}

func (p Philosopher) eat() {
	defer eatGroup.Done()
	for j := 0; j < 3; j++ {
		p.leftFork.Lock()
		p.rightFork.Lock()

		say("eating", p.id)
		time.Sleep(time.Second)

		p.rightFork.Unlock()
		p.leftFork.Unlock()

		say("finished eating", p.id)
		time.Sleep(time.Second)
	}

}

func say(action string, id int) {
	fmt.Printf("Philosopher #%d is %s\n", id+1, action)
}

func main() {
	count := 5

	// Create forks
	forks := make([]*fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(fork)
	}

	philosophers := make([]*Philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &Philosopher{
			id: i, leftFork: forks[i], rightFork: forks[(i+1)%count]}
		eatGroup.Add(1)
		go philosophers[i].eat()

	}
	eatGroup.Wait()

}

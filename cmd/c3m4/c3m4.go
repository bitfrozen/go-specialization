package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Chopstick struct {
	id        int
	timesUsed int
}

type Ticket struct {
	chopsticks [2]*Chopstick
}

type Philosopher struct {
	name string
}

func (p Philosopher) eat(receive <-chan *Ticket, giveBack chan<- *Ticket, done chan<- bool) {
	for i := 0; i < 3; i++ {
		// request chopsticks (aka permission)
		ticket := <-receive

		fmt.Printf("%s starting to eat with chopsticks %d and %d\n", p.name, ticket.chopsticks[0].id, ticket.chopsticks[1].id)
		// Adding slight delay to check how many philosophers eat at the same time
		time.Sleep(time.Duration(rand.Intn(1e1)) * time.Millisecond)
		ticket.chopsticks[0].timesUsed++
		ticket.chopsticks[1].timesUsed++
		fmt.Printf("%s finishing eating with chopsticks %d and %d\n", p.name, ticket.chopsticks[0].id, ticket.chopsticks[1].id)

		// give back chopsticks (aka report finished usage)
		giveBack <- ticket
	}

	// report finished meals
	fmt.Printf("%s done with meals\n", p.name)
	done <- true
}

func host(wg *sync.WaitGroup, chopsticks chan *Chopstick, philosophers []Philosopher) {
	// make this channel buffered to avoid locking up because philosophers can't report done status
	doneEating := make(chan bool, len(philosophers))

	// tickets management channels
	sendingSticks := make(chan *Ticket, 2)
	receivingSticks := make(chan *Ticket, 2)
	tray := make(chan *Ticket, 2)

	for _, p := range philosophers {
		go p.eat(sendingSticks, receivingSticks, doneEating)
	}

	for i := 0; i < 2; i++ {
		ticket := new(Ticket)
		shuffleSticks(chopsticks)
		ticket.chopsticks[0] = <-chopsticks
		shuffleSticks(chopsticks)
		ticket.chopsticks[1] = <-chopsticks
		tray <- ticket
	}

	for numberOfFedPhilosophers := 0; numberOfFedPhilosophers < 5; {
		select {

		case ticket := <-tray:
			sendingSticks <- ticket

		case ticket := <-receivingSticks:
			chopsticks <- ticket.chopsticks[0]
			chopsticks <- ticket.chopsticks[1]

			shuffleSticks(chopsticks)
			ticket.chopsticks[0] = <-chopsticks
			shuffleSticks(chopsticks)
			ticket.chopsticks[1] = <-chopsticks
			tray <- ticket

		case <-doneEating:
			numberOfFedPhilosophers++
		}
	}

	close(tray)
	close(sendingSticks)
	close(receivingSticks)
	close(doneEating)

	// gather back all chopsticks (for auditing and debug)
	for ticket := range sendingSticks {
		chopsticks <- ticket.chopsticks[0]
		chopsticks <- ticket.chopsticks[1]
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	philosophers := make([]Philosopher, 5)
	for i := range philosophers {
		// philosophers with letter names are easier to debug
		philosophers[i] = Philosopher{name: string(rune(65 + i))}
	}

	chopsticks := make(chan *Chopstick, 5)
	for i := 0; i < 5; i++ {
		c := Chopstick{id: i + 1}
		chopsticks <- &c
	}

	wg.Add(1)
	go host(&wg, chopsticks, philosophers)
	wg.Wait()

	close(chopsticks)
	// print chopstick statistics for auditing
	for stick := range chopsticks {
		if stick != nil {
			fmt.Printf("-- chopstick: %d used %d times\n", stick.id, stick.timesUsed)
		}
	}
}

func shuffleSticks(sticks chan *Chopstick) {
	times := rand.Intn(cap(sticks)-1) + 1

	for i := 0; i < times; i++ {
		val := <-sticks
		sticks <- val
	}
}

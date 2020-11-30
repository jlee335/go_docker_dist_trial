package main

import (
	"fmt"
	"time"
)

type Job struct {
	calc      int
	uid       int
	terminate bool
}

func main() {
	sender := make(chan Job, 100)
	reciever := make(chan Job, 100)
	//var wg sync.WaitGroup
	//wg.Add(2)

	go worker(sender, reciever)
	go worker(sender, reciever)

	for uid := 0; uid < 100; uid++ {
		sender <- Job{calc: uid, uid: uid}
	}

	for res := 0; res < 100; res++ {
		fmt.Println(<-reciever)
	}

}

// This worker calculates x^2 and returns to channel
func worker(jobs chan Job, result chan Job) {
	for n := range jobs {
		x := n.calc
		y := x * x * x
		time.Sleep(time.Millisecond * 1000)
		n.calc = y
		result <- n
	}
}

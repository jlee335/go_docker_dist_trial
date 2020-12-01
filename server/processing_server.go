package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Job struct {
	calc      int
	uid       int
	terminate bool
}

//Function that initializes workers
func init_workers(sender chan Job, reciever chan Job) {
	go worker(sender, reciever)
	go worker(sender, reciever)
	go worker(sender, reciever)
	go worker(sender, reciever)
}

func main() {
	sender := make(chan Job, 100)
	reciever := make(chan Job, 100)
	//var wg sync.WaitGroup
	//wg.Add(2)
	init_workers(sender, reciever)

	fmt.Printf("Starting server at port 8081\n")

	http.HandleFunc("/", calcHandler)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "pong from processing server")

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

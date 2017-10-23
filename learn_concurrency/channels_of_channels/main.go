package main

import (
	"time"
)

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}

	return
}

func handle(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}

func main() {
	ClientRequests := make(chan *Request, 1)
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}

	ClientRequests <- request
	go handle(ClientRequests)
	time.Sleep(time.Millisecond * 100)
}

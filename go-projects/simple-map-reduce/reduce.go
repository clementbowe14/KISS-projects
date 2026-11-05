package main

import(
	"sync"
)

type Reducer struct {
	threadCount int
	wg *sync.WaitGroup
	wordMap map[string]string
	mu sync.Mutex
}

func (r *Reducer) worker(workerID int, resultChannel chan map[string]int) {
	defer r.wg.Done()
}
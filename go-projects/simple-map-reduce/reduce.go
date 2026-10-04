package main

import(
	"sync"
)

type Reduce struct {
	wordMap map[string]string
	mu sync.Mutex
}
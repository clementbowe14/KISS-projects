package main

import (
	"log"
	"os"
	"sync"
)

type Coordinator struct {
	mapper      *Map
	numWorkers  int
	logger      *log.Logger
}

func NewCoordinator(numWorkers int) *Coordinator {
	mapper := &Map{
		NumWorkers: numWorkers,
		wg:         &sync.WaitGroup{},
		Logger:     log.New(os.Stdout, "MapWorker: ", log.LstdFlags),
	}

	return &Coordinator{
		mapper: mapper,
		logger: log.New(os.Stdout, "Coordinator: ", log.LstdFlags),
	}
}

func (c *Coordinator) Start() {
	c.logger.Println("Starting the MapReduce Coordinator...")
	mapperFiles, err := os.ReadDir("./files")
	if err != nil {
		c.logger.Fatalf("Error reading files directory: %v", err)
	}
	c.mapper.SetFiles(mapperFiles)
	c.mapper.Run()
	c.logger.Println("MapReduce Coordinator finished.")
}

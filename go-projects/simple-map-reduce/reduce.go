package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
)

type Reducer struct {
	NumWorkers int
	wg         *sync.WaitGroup
	logger     *log.Logger
}

func (r *Reducer) Run() {
	for i := 0; i < r.NumWorkers; i++ {
		r.wg.Add(1)
		go r.worker(i)
	}

	r.wg.Wait()

}

func (r *Reducer) worker(workerID int) {
	defer r.wg.Done()
	files, err := r.getFiles(workerID)
	wg := &sync.WaitGroup{}
	mappedFilesDir := "./mapped-files"
	wordFrequency := make(map[string]int)
	mu := &sync.Mutex{}

	if err != nil {
		log.Printf("Worker %d: Error retrieving files: %v", workerID, err)
		return
	}

	for _, file := range files {
		wg.Add(1)
		go func(file os.DirEntry, wordFrequency map[string]int) {
			defer wg.Done()
			mappedFilePath := mappedFilesDir + "/" + file.Name()
			mappedFile, err := os.OpenFile(mappedFilePath, os.O_RDONLY, 0644)
			if err != nil {
				log.Printf("Worker %d: Error opening file %s: %v", workerID, mappedFilePath, err)
				return
			}

			if err != nil {
				log.Printf("Worker %d: Error opening file %s: %v", workerID, mappedFilePath, err)
				return
			}

			defer mappedFile.Close()
			mapperResult := MapperResultFile{}
			decoder := json.NewDecoder(mappedFile)
			if err := decoder.Decode(&mapperResult); err != nil {
				log.Printf("Worker %d: Error decoding JSON from file %s: %v", workerID, mappedFilePath, err)
				return
			}

			for _, wordRecord := range mapperResult.WordEntries {
				mu.Lock()
				wordFrequency[wordRecord.Word] += wordRecord.Count
				mu.Unlock()
			}

		}(file, wordFrequency)
	}

	wg.Wait()

	reducerOutputFilePath := fmt.Sprintf("./reduced-files/reducer-%d.json", workerID)
	reducerOutputFile, err := os.Create(reducerOutputFilePath)
	if err != nil {
		log.Printf("Worker %d: Error creating reducer output file %s: %v", workerID, reducerOutputFilePath, err)
		return
	}

	defer reducerOutputFile.Close()
	
	reducerResult := MapperResultFile{}
	for word, count := range wordFrequency {
		reducerResult.WordEntries = append(reducerResult.WordEntries, WordRecord{
			Word:  word,
			Count: count,
		})
	}

	encoder := json.NewEncoder(reducerOutputFile)
	if err := encoder.Encode(&reducerResult); err != nil {
		log.Printf("Worker %d: Error encoding JSON to file %s: %v", workerID, reducerOutputFilePath, err)
		return
	}

	log.Printf("Worker %d: Successfully reduced and wrote output to %s", workerID, reducerOutputFilePath)

}

func (r *Reducer) getFiles(workerID int) ([]os.DirEntry, error) {
	// Implement logic to retrieve files for the given workerID
	// This is a placeholder implementation; you should replace it with actual file retrieval logic.
	files, err := os.ReadDir("./mapped-files")
	if err != nil {
		r.logger.Printf("Worker %d: Error reading directory: %v", workerID, err)
		return nil, fmt.Errorf("error reading directory: %v", err)
	}
	var reducerFiles []os.DirEntry
	reducerPartionRegex := regexp.MustCompile(fmt.Sprintf(`reducer-%d.json`, workerID))
	for _, file := range files {
		if !file.IsDir() && reducerPartionRegex.MatchString(file.Name()) {
			// You can add logic here to filter files based on workerID if needed
			// For example, you could assign files to workers based on their index or name
			fmt.Printf("Worker %d: Found file %s\n", workerID, file.Name())
			reducerFiles = append(reducerFiles, file)
		}
	}
	return reducerFiles, nil
}

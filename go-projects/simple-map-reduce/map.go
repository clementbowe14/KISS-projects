package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"strings"
	"sync"
)

type Map struct {
	threadCount int
	wg          *sync.WaitGroup
	Logger      *log.Logger
	NumWorkers  int
	numTasks    int
	files       []os.DirEntry
	fileChannel chan string
}

type WordRecord struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type MapperResultFile struct {
	WordEntries []WordRecord `json:"word_entries"`
}

func (m *Map) Run() {
	for i := 0; i < m.NumWorkers; i++ {
		m.wg.Add(1)
		go m.worker(i)
	}
	for _, file := range m.files {
		if !file.IsDir() {
			filePath := file.Name()
			m.fileChannel <- filePath
		}
	}

	close(m.fileChannel)
	m.wg.Wait()
}

func (m *Map) worker(workerID int) {
	defer m.wg.Done()

	for filePath := range m.fileChannel {
		m.Logger.Printf("Worker %d processing file: %s", workerID, filePath)
		finalFilePath := "./files/" + filePath
		file, err := os.Open(finalFilePath)
		if err != nil {
			m.Logger.Printf("Worker %d: Error opening file %s: %v", workerID, filePath, err)
			continue
		}

		scanner := bufio.NewScanner(file)
		wordMap := make(map[string]int)

		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Fields(line)
			for _, word := range words {
				wordMap[word]++
			}
		}

		if err := scanner.Err(); err != nil {
			m.Logger.Printf("Worker %d: Error reading file %s: %v", workerID, filePath, err)
		}

		file.Close()

		reducerFilePath := "./mapped-files/"

		for word, count := range wordMap {
			reducerFileName := fmt.Sprintf("%smapper-%d-reducer-%d.json", reducerFilePath, workerID, m.GetReducerPartition(word))
			reducerFile, err := os.OpenFile(reducerFileName, os.O_CREATE|os.O_RDWR, 0644)

			if err != nil {
				m.Logger.Printf("Worker %d: Error opening reducer file %s: %v", workerID, reducerFileName, err)
				continue
			}

			info, err := reducerFile.Stat()
			if err != nil {
				m.Logger.Printf("Worker %d: Error getting file info for %s: %v", workerID, reducerFileName, err)
				reducerFile.Close()
				continue
			}

			if info.Size() == 0 {
				mapperFile := MapperResultFile{
					WordEntries: []WordRecord{},
				}
				mapperFile.WordEntries = append(mapperFile.WordEntries, WordRecord{Word: word, Count: count})
				data, err := json.Marshal(mapperFile)
				if err != nil {
					m.Logger.Printf("Worker %d: Error marshalling JSON for reducer file %s: %v", workerID, reducerFileName, err)
					reducerFile.Close()
					continue
				}

				err = os.WriteFile(reducerFileName, data, 0644)
				if err != nil {
					m.Logger.Printf("Worker %d: Error writing to reducer file %s: %v", workerID, reducerFileName, err)
					reducerFile.Close()
					continue
				}

			} else {
				mapperFile := MapperResultFile{}
				decoder := json.NewDecoder(reducerFile)
				decoder.Decode(&mapperFile)
				mapperFile.WordEntries = append(mapperFile.WordEntries, WordRecord{Word: word, Count: count})
				data, err := json.Marshal(mapperFile)
				if err != nil {
					m.Logger.Printf("Worker %d: Error marshalling JSON for reducer file %s: %v", workerID, reducerFileName, err)
					reducerFile.Close()
					continue
				}

				err = os.WriteFile(reducerFileName, data, 0644)
				if err != nil {
					m.Logger.Printf("Worker %d: Error writing to reducer file %s: %v", workerID, reducerFileName, err)
					reducerFile.Close()
					continue
				}
				reducerFile.Close()
			}

		}
	}
}

func (m *Map) GetReducerPartition(word string) int {
	hash := fnv.New64a()
	hash.Write([]byte(word))
	hashValue := hash.Sum64()
	if int(hashValue) % m.NumWorkers < 0 {
		return -1 * (int(hashValue) % m.NumWorkers)
	}
	return int(hashValue) % m.NumWorkers
}

func (m *Map) SetFiles(files []os.DirEntry) {
	m.files = files
	m.fileChannel = make(chan string, m.NumWorkers)
}

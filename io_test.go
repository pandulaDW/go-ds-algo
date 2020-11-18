package main

import (
	"path/filepath"
	"sync"
	"testing"

	"algos.com/main/iostuff"
)

var writePath, readPath string
var data []string

func init() {
	writePath, _ = filepath.Abs(filepath.Join("data", "data_write.txt"))
	readPath, _ = filepath.Abs(filepath.Join("data", "data_read.txt"))
	data = []string{
		"A bad workman always blames his tools.",
		"A bird in hand is worth two in the bush.",
		"Absence makes the heart grow fonder.",
	}
}

func BenchmarkSeqIOTest(t *testing.B) {
	for i := 0; i < t.N; i++ {
		iostuff.WriteFileStream(data, writePath)
		content := iostuff.ReadFileStream(readPath)
		_ = content
	}
}

func BenchmarkConcurrentTest(t *testing.B) {
	var wg sync.WaitGroup
	ch := make(chan string)

	for i := 0; i < t.N; i++ {
		wg.Add(2)

		go func() {
			iostuff.WriteFileStream(data, writePath)
			defer wg.Done()
		}()

		go func() {
			ch <- iostuff.ReadFileStream(readPath)
			defer wg.Done()
		}()

		content := <-ch
		_ = content

		wg.Wait()
	}
}

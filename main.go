package main

import (
	"path/filepath"

	"algos.com/main/iostuff"
)

func main() {
	inputPath, _ := filepath.Abs(filepath.Join("data", "data_read.txt"))

	iostuff.BufferedIORead(inputPath)
}

package main

import (
	"fmt"
	"path/filepath"

	"algos.com/main/iostuff"
)

func main() {
	path, _ := filepath.Abs(filepath.Join("data", "data_write.txt"))
	data := []string{
		"A bad workman always blames his tools.",
		"A bird in hand is worth two in the bush.",
		"Absence makes the heart grow fonder.",
	}

	iostuff.WriteFileStream(data, path)

	path, _ = filepath.Abs(filepath.Join("data", "data_read.txt"))
	content := iostuff.ReadFileStream(path)
	fmt.Println(content)
}

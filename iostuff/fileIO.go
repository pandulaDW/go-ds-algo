package iostuff

import (
	"fmt"
	"io"
	"os"
)

// Type os.File represents a file on the local system. It implements
// both io.Reader and io.Writer and, therefore, can be used in any streaming IO contexts

// WriteFileStream writes data to a given file. If the file doesn't exists, it will create a file first
func WriteFileStream(data []string, path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, p := range data {
		n, err := file.Write([]byte(p))
		_, err = file.Write([]byte{'\n'})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println("file write done")
}

// ReadFileStream reads data from a given file and return the content as a string
func ReadFileStream(path string) string {
	file, err := os.OpenFile(path, os.O_RDONLY, 0444)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var content []byte
	line := make([]byte, 20) // reading 20 bytes at a time

	for {
		n, err := file.Read(line)
		if err == io.EOF {
			break
		}
		content = append(content, line[:n]...)
	}

	return string(content)
}

package iostuff

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// Type os.File represents a file on the local system. It implements
// both io.Reader and io.Writer and, therefore, can be used in any streaming IO contexts

// io.Copy() makes it easy to stream data from a source reader to a target writer.
// It abstracts out the for-loop pattern (we’ve seen so far) and properly handle io.EOF and byte counts

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

// WriteStdOut writes data to standard output
func WriteStdOut(content string) {
	data := strings.Split(content, "\n")

	for _, val := range data {
		n, err := os.Stdout.Write([]byte(val))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(val) {
			fmt.Println("Failed to write data")
			os.Exit(1)
		}
	}
}

// CopyStream reads data from a file and write to the another file
func CopyStream(inputPath string, outputPath string) {
	inputFile, err := os.OpenFile(inputPath, os.O_RDONLY, 0444)
	outputFile, err := os.Create(outputPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n, err := io.Copy(outputFile, inputFile)

	if err == nil {
		fmt.Printf("\nSuccessfully copied %d bytes\n", n)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	defer inputFile.Close()
	defer outputFile.Close()
}

// PipeReaderWriter creates a synchronous in-memory pipe.
// It can be used to connect code expecting an io.Reader with code expecting an io.Writer.
func PipeReaderWriter(content string) {
	byteBuffer := new(bytes.Buffer)
	byteBuffer.WriteString(content)

	piper, pipew := io.Pipe()

	// write in writer end of pipe
	go func() {
		defer pipew.Close()
		io.Copy(pipew, byteBuffer)
	}()

	// read from reader end of the pipe
	io.Copy(os.Stdout, piper)
}

// BufferedIORead will read the content using a buffer
func BufferedIORead(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Println(line)
	}
}

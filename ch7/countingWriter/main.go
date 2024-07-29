package main

import (
	"fmt"
	"io"
)

// countingWriter is a struct that embeds an io.Writer and tracks the count of bytes written.
type CountingWriterStruct struct {
	writer io.Writer
	count  *int64
}

// Write method writes data to the underlying writer and updates the count.
func (cw *CountingWriterStruct) Write(p []byte) (int, error) {
	n, err := cw.writer.Write(p)
	*cw.count += int64(n)
	return n, err
}

// CountingWriter returns a new Writer that wraps the original, and a pointer to an int64 variable
// that contains the number of bytes written to the new Writer.
func CountingWriter(w io.Writer) (CountingWriterStruct, *int64) {
	count := int64(0)
	cw := CountingWriterStruct{
		writer: w,
		count:  &count,
	}
	return cw, &count
}

// Example usage
func main() {
	var w io.Writer = &exampleWriter{} // Replace exampleWriter with your actual writer

	writer, count := CountingWriter(w)

	data := []byte("Hello, World!")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Bytes written: %d, Total count: %d\n", n, *count)
}

// exampleWriter is a simple implementation of io.Writer for demonstration purposes.
type exampleWriter struct{}

func (ew *exampleWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

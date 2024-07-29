package main

import (
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	r io.Reader
	n int64
}

func (lr *limitReader) Read(p []byte) (int, error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}

	fmt.Println(len(p))

	if int(lr.n) < len(p) {
		p = p[:lr.n]
	}

	n, err := lr.r.Read(p)
	lr.n -= int64(n)

	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	lr := limitReader{r, n}

	return &lr
}

func main() {
	r := strings.NewReader("Hello, Reader!")

	lr := LimitReader(r, 5)

	b := make([]byte, 8)
	for {
		n, err := lr.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	contents string
	pos      int
}

func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos <= len(m.contents) {
		n := copy(p, m.contents[m.pos:m.pos+1])
		m.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main1() {
	mySlowReaderInstance := &MySlowReader{
		contents: "Hello World!",
	}
	out, err := io.ReadAll(mySlowReaderInstance)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output: %s\n", out)
}

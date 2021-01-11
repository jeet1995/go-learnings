package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var wc WriterCloserNew = NewBufferedWriterCloser()
	wc.Write([]byte("Hello YouTube listeners, this is a test"))
	wc.Close()

	// type conversion
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	// type conversion error
	// bwc := wc.(io.Reader)

	reader, ok := wc.(io.Reader)

	if ok {
		fmt.Println(reader)
	} else {
		fmt.Println("conversion failed")
	}

	// everything can be cast to an empty interface
	// even primitives
	var sumObject interface{} = NewBufferedWriterCloser()

	reader, ok = sumObject.(io.Reader)

	if ok {
		fmt.Println(reader)
	} else {
		fmt.Println("conversion failed")
	}
}

type WriterNew interface {
	Write([]byte) (int, error)
}

type CloserNew interface {
	Close() error
}

// embedding an interface
type WriterCloserNew interface {
	WriterNew
	CloserNew
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)

	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)

	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

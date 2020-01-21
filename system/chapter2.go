package main

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/jnst/x-go/system/server"
)

// Chapter2 represents chapter-2 of go system programming book.
type Chapter2 struct{}

func (c Chapter2) Q1() {
	if _, err := fmt.Fprintf(os.Stdout, "Chapter-%d Q%f\n", 1, 1.0); err != nil {
		panic(err)
	}
}

func (c Chapter2) Q2() {
	w := csv.NewWriter(os.Stdout)
	if err := w.Write([]string{"a", "b", "c"}); err != nil {
		panic(err)
	}
	w.Flush()
}

func (c Chapter2) Q3() {
	s := server.NewServer(c.gzipHandler)
	s.Run()
}

func (c Chapter2) gzipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}

	writer := io.MultiWriter(w, os.Stdout)
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(source); err != nil {
		panic(err)
	}

	gzipWriter := gzip.NewWriter(w)
	if err := gzipWriter.Flush(); err != nil {
		panic(err)
	}
	if err := gzipWriter.Close(); err != nil {
		panic(err)
	}
}

func (c Chapter2) CreateFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte("os.File example\n")) // io.Writer interface
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	// clean up
	info, err := os.Stat(name)
	if err == nil {
		fmt.Printf("filename: %s was found. This file being deleted.\n", info.Name())
		err = os.Remove(name)
		if err != nil {
			panic(err)
		}
	}
}

func (c Chapter2) Print(message string) {
	if _, err := os.Stdout.Write([]byte(message)); err != nil { // io.Writer interface
		panic(err)
	}
}

func (c Chapter2) PrintByBuffer(message string) {
	var buf bytes.Buffer
	if _, err := buf.Write([]byte(message)); err != nil { // io.Writer interface
		panic(err)
	}
	if _, err := io.WriteString(&buf, message); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

func (c Chapter2) Connect() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	if _, err := io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"); err != nil {
		panic(err)
	}
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	if err := req.Write(conn); err != nil {
		panic(err)
	}
}

func main() {
	c := Chapter2{}
	fmt.Println("--- Chapter2 Q1 ---")
	c.Q1()
	fmt.Println("--- Chapter2 Q2 ---")
	c.Q2()
	fmt.Println("--- Chapter2 Q3 ---")
	c.Q3()
}

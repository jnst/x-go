package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// PrintFileByScanner reads specified file and print to stdout by using scanner
func PrintFileByScanner(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

// PrintFileByIOUtil reads specified file and print to stdout by using ioutil
func PrintFileByIOUtil(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// List reads current directory files then print to stdout
func List() {
	name, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(name)
	names, err := f.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range names {
		fmt.Println(s)
	}
}

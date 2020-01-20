package main

import (
	"io"
	"io/ioutil"
	"os"
)

// Chapter2 represents chapter-3 of go system programming book.
type Chapter3 struct {}

func (c Chapter3) Q1(dstPath, srcPath string) {
	//copyWithIOUtil(dstPath, srcPath)
	copyWithOS(dstPath, srcPath)
}

func copyWithIOUtil(dstPath, srcPath string) {
	b, err := ioutil.ReadFile(srcPath)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(dstPath, b, 0644); err != nil {
		panic(err)
	}
}

func copyWithOS(dstPath, srcPath string) {
	rf, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	defer close(rf)

	wf, err := os.OpenFile(dstPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer close(wf)

	if _, err := io.Copy(wf, rf); err != nil {
		panic(err)
	}
}

func close(f *os.File) {
	if err := f.Close(); err != nil {
		panic(err)
	}
}

func (c Chapter3) Q2() {

}

func (c Chapter3) Q3() {

}

func (c Chapter3) Q4() {

}

func (c Chapter3) Q5() {

}

func (c Chapter3) Q6() {

}

func main() {
	c := Chapter3{}
	c.Q1("testdata/new.txt", "testdata/old.txt")
}

package main

import (
	"archive/zip"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/jnst/x-go/system/server"
)

// Chapter2 represents chapter-3 of go system programming book.
type Chapter3 struct{}

// Q1 copies the file.
func (c Chapter3) Q1(dstPath, srcPath string) {
	copyWithIOUtil(dstPath, srcPath)
	copyWithOS(dstPath, srcPath)
}

func copyWithIOUtil(dstPath, srcPath string) {
	b, err := ioutil.ReadFile(srcPath)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(dstPath, b, 0o600); err != nil {
		panic(err)
	}
}

func copyWithOS(dstPath, srcPath string) {
	rf, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	defer rf.Close()

	wf, err := os.OpenFile(dstPath, os.O_RDWR, 0o600)
	if err != nil {
		panic(err)
	}
	defer wf.Close()

	if _, err := io.Copy(wf, rf); err != nil {
		panic(err)
	}
}

// Q2 creates an appropriately sized file for testing.
func (c Chapter3) Q2(path string, size int64) {
	b := make([]byte, size)
	r := rand.Reader

	if _, err := r.Read(b); err != nil {
		panic(err)
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := io.CopyN(f, r, size); err != nil {
		panic(err)
	}

	verify(path, size)
}

func verify(path string, size int64) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var (
		w   io.Writer
		msg string
	)

	if int64(len(b)) == size {
		w = os.Stdout
		msg = "success"
	} else {
		w = os.Stderr
		msg = "failure"
	}

	if _, err := w.Write([]byte(msg)); err != nil {
		panic(err)
	}
}

// Q3 writes to zip file.
func (c Chapter3) Q3() {
	zipFile, err := os.Create("testdata/archive.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		if err := zipWriter.Close(); err != nil {
			panic(err)
		}
	}()

	w, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadFile("testdata/file.txt")
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		panic(err)
	}
}

// Q4 downloads from web server.
func (c Chapter3) Q4() {
	s := server.NewServer(c.zipHandler)
	s.Run()
}

func (c Chapter3) zipHandler(w http.ResponseWriter, r *http.Request) {
	const FILENAME = "ascii_sample.zip"

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", FILENAME))

	zipWriter := zip.NewWriter(w)

	for i := 0; i < 3; i++ {
		writer, err := zipWriter.Create(fmt.Sprintf("sample-%d.txt", i))
		if err != nil {
			panic(err)
		}

		if _, err := writer.Write([]byte("this is text file " + strconv.Itoa(i))); err != nil {
			panic(err)
		}
	}

	if err := zipWriter.Flush(); err != nil {
		panic(err)
	}

	if err := zipWriter.Close(); err != nil {
		panic(err)
	}
}

// Q5 implements io.CopyN.
func (c Chapter3) Q5(dst io.Writer, src io.Reader, n int64) (int64, error) {
	r := io.LimitReader(src, n)

	written, err := io.Copy(dst, r)
	if err != nil {
		return written, errors.Wrap(err, err.Error())
	}

	return written, nil
}

// Q6 prints "ASCII".
func (c Chapter3) Q6() {
	computer := strings.NewReader("COMPUTER")
	system := strings.NewReader("SYSTEM")
	programming := strings.NewReader("PROGRAMMING")

	var stream io.Reader

	// write code here
	a := io.NewSectionReader(programming, 5, 1)
	s := io.NewSectionReader(system, 0, 1)
	cc := io.NewSectionReader(computer, 0, 1)
	i1 := io.NewSectionReader(programming, 8, 1)
	i2 := io.NewSectionReader(programming, 8, 1)

	stream = io.MultiReader(a, s, cc, i1, i2)

	if _, err := io.Copy(os.Stdout, stream); err != nil {
		panic(err)
	}
}

/*func main() {
	c := Chapter3{}
	c.Q1("testdata/new.txt", "testdata/old.txt")
	c.Q2("testdata/binary.txt", 1024)
	c.Q3()
	c.Q4()
	c.Q6()
}*/

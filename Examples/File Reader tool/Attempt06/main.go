package main

import (
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
	"os"
	"io"
	"bytes"
)

var file = flag.String("file", "/var/tmp/text", "source file")
var alines =flag.Int("alines", 0, "lines to read, starting from beginning")
var zlines = flag.Int("zlines", 0, "lines to read, starting from end")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// reads file in chunks until line count satisfied
func file_reader(file string, size int64) chan []byte {
	fileout := make(chan []byte)
	f, err := os.Open(file)
	check(err)
  chunk_size := 500
	data := make([]byte, chunk_size)
	go func() {
		for i := int(size); i > 0; i -= chunk_size {
			if i >= chunk_size {
				fmt.Printf("***reading %v bytes, starting at offset %v***\n", len(data), i - chunk_size)
				count, err := f.ReadAt(data, int64(i - chunk_size))
				check(err)
				fileout <- count
				close(fileout)
			} else {
				fmt.Printf("***less than %v bytes left, printing remaining %v from beginning***\n", len(data), i)
				remainder := make([]byte, i)
				count, err := f.Read(remainder)
				check(err)
				fileout <- count
				close(fileout)
			}
		}
	}()
	return fileout
}

//will find lines from data received from filereader and return them in array
func line_finder(data string, lines int) []string {
  array := make([]byte, lines)
}

//counts lines so that line numbers will be correct
func line_counter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

//final print of requested data with line numbers
func request_printer(answer []byte) {
	for i := 0; i < len(answer); i++ {
		fmt.Println(i + 1, ":", answer[i])
	}
}



//prints request as received from flags, returns file size
func printrequest(file string, a,z int) int64 {
	stat, err := os.Stat(file)
	check(err)
	fmt.Printf("Filename: '%s'|Size: %d bytes|Modified: %v\n", stat.Name(),stat.Size(),stat.ModTime())
	return stat.Size()
}


func main() {

	flag.Parse()

	size := printrequest(*file, *alines, *zlines)
	file_reader(*file, size)
}
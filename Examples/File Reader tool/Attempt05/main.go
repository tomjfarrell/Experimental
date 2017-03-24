package main

import (
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
	"os"
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
func filereader(file string, size int64) {
  f, err := os.Open(file)
	check(err)
  chunk_size := 500
	data := make([]byte, chunk_size)
	for i := int(size); i > 0; i-= chunk_size {
		if i >= chunk_size {
			fmt.Printf("***reading %v bytes, starting at offset %v***\n",len(data),i- chunk_size)
			count, err := f.ReadAt(data,int64(i- chunk_size))
			check(err)
			contents := strings.Split(string(data[:count]), "\n")
			length := len(contents)-1
			for i := 0; i < length; i++ {
				fmt.Println(i + 1, ":", contents[i])
			}
			fmt.Printf("read %d bytes: %q\n", count, data[:count])
		} else {
			fmt.Printf("***less than %v bytes left, printing remaining %v from beginning***\n",len(data),i)
			remainder := make([]byte,i)
			count, err := f.Read(remainder)
			check(err)
			contents := strings.Split(string(remainder[:count]), "\n")
			length := len(contents)-1
			for i := 0; i < length; i++ {
				fmt.Println(i + 1, ":", contents[i])
			}
			fmt.Printf("read %d bytes: %q\n", count, remainder[:count])

		}
	}
}

//will find lines from data received from filereader and return them in array
func linefinder(data string, lines int) []byte {
  array := make([]byte, lines)
}

//final print of requested data with line numbers
func request_printer(answer []byte) {
	for i := 0; i < len(answer); i++ {
		fmt.Println(i + 1, ":", answer[i])
	}
}

//prints request as received from flags
func printrequest(file string, a,z int) int64 {
	stat, err := os.Stat(file)
	check(err)
	fmt.Printf("Filename: '%s'|Size: %d bytes|Modified: %v\n", stat.Name(),stat.Size(),stat.ModTime())
	return stat.Size()
}


func main() {

	flag.Parse()

	size := printrequest(*file,*alines,*zlines)
  filereader(*file,size)
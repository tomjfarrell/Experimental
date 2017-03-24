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

func filereader(file string) []byte {
	dat, err := ioutil.ReadFile(file)
	check(err)
	return dat
}
func filereader2(file string, size int64) {
  f, err := os.Open(file)
	check(err)
  chunksize := 500
	data := make([]byte, chunksize)
	for i := int(size); i > 0; i-=chunksize {
		if i >= 500 {
			fmt.Printf("***reading %v bytes, starting at offset %v***\n",len(data),i-chunksize)
			count, err := f.ReadAt(data,int64(i-chunksize))
			check(err)
			fmt.Println(string(data[:count]))
		} else {
			fmt.Printf("***less than %v bytes left, printing remaining %v from beginning***\n",len(data),i)
			remainder := make([]byte,i)
			count, err := f.Read(remainder)
			check(err)
			fmt.Println(string(remainder[:count]))
		}
	}
}

func printrequest(file string, a,z int) int64 {
	stat, err := os.Stat(file)
	check(err)
	fmt.Printf("Filename: '%s'|Size: %d bytes|Modified: %v\n", stat.Name(),stat.Size(),stat.ModTime())
	//fmt.Printf("file: %s, length of file %d\n", file, l)
	if (a == 0) && (z == 0) {
		fmt.Println("Need --alines and/or --zlines flag(s)")
	}
	return stat.Size()
}


func main() {

	flag.Parse()

	data := filereader(*file)
	contents := strings.Split(string(data), "\n")
	length := len(contents)-1
	size := printrequest(*file,*alines,*zlines)
  filereader2(*file,size)

	if (length < (*alines + *zlines)) {
		fmt.Println("Not enough contents to fulfill request, printing entire file.")
		fmt.Printf("There will be an overlap of %d line(s)\n",(*alines+*zlines-length))
		for i := 0; i < length; i++ {
			fmt.Println(i+1, ":", contents[i])
		}
	} else {
		if *alines > 0 {
			fmt.Printf("reading first %d lines:\n", *alines)
			for i := 0; i < *alines; i++ {
				fmt.Println(i+1, ":", contents[i])
			}
		}
		if *zlines > 0 {
			fmt.Printf("reading last %d lines:\n", *zlines)
			for i := length-int(*zlines); i < length; i++ {
				fmt.Println(i+1,":",contents[i])
			}
		}
	}
}

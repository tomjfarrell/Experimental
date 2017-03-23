package main

import (
	"io/ioutil"
	"os"
	"flag"
	"fmt"
)

var fileflag = flag.String("file", "/var/tmp/text.rtf", "source file")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type file string

func read(file) []byte {
	dat, err := ioutil.ReadFile(file)
	check(err)
	return dat
}

func open(file) *os.File {
	f, err := os.Open(file)
	check(err)
	return f
}

func main() {

	flag.Parse()
	file := *fileflag

  fmt.Println(file)

	dat := read(file)

	f := open(file)

	f.Close()

	//// Read some bytes from the beginning of the file.
	//// Allow up to 5 to be read but also note how many
	//// actually were read.
	//b1 := make([]byte, 5)
	//n1, err := f.Read(b1)
	//check(err)
	//fmt.Printf("%d bytes: %s\n", n1, string(b1))
	//
	//// You can also `Seek` to a known location in the file
	//// and `Read` from there.
	//o2, err := f.Seek(6, 0)
	//check(err)
	//b2 := make([]byte, 2)
	//n2, err := f.Read(b2)
	//check(err)
	//fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	//
	//// The `io` package provides some functions that may
	//// be helpful for file reading. For example, reads
	//// like the ones above can be more robustly
	//// implemented with `ReadAtLeast`.
	//o3, err := f.Seek(6, 0)
	//check(err)
	//b3 := make([]byte, 2)
	//n3, err := io.ReadAtLeast(f, b3, 2)
	//check(err)
	//fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	//
	//// There is no built-in rewind, but `Seek(0, 0)`
	//// accomplishes this.
	//_, err = f.Seek(0, 0)
	//check(err)
	//
	//// The `bufio` package implements a buffered
	//// reader that may be useful both for its efficiency
	//// with many small reads and because of the additional
	//// reading methods it provides.
	//r4 := bufio.NewReader(f)
	//b4, err := r4.Peek(5)
	//check(err)
	//fmt.Printf("5 bytes: %s\n", string(b4))

}

package main

import (
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
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

func printrequest(file string, l,a,z int) {
	fmt.Printf("file: %s, length of file %d\n", file, l)
	if (a == 0) && (z == 0) {
		fmt.Println("Need --alines and/or --zlines flag(s)")
	}
}


func main() {

	flag.Parse()

	data := filereader(*file)
	contents := strings.Split(string(data), "\n")

	printrequest(*file,len(contents),*alines,*zlines)

	if (len(contents) < *alines) && (len(contents) < *zlines) {
		fmt.Println("Not enough contents to fulfill request, printing entire file.")
		for i := 0; i < len(contents)-1; i++ {
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
			fmt.Printf("reading last %d lines:\n", *alines)
			for i := len(contents)-1-int(*zlines); i < len(contents)-1; i++ {
				fmt.Println(i+1,":",contents[i])
			}
		}
	}
}

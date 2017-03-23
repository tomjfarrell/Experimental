package main

import (
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
)

var file = flag.String("file", "/var/tmp/text", "source file")
var alines =flag.Int("alines", 10, "lines to read, starting from beginning")
var zlines = flag.Int("zlines", 10, "lines to read, starting from end")

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

func main() {

	flag.Parse()

	fmt.Printf("file: %s, counting back %d lines\n",*file,*lines)

	data := filereader(*file)
	contents := strings.Split(string(data), "\n")
	fmt.Printf("lenth of file: %d\n",len(contents)-1)
	if len(contents) < 10 {
		fmt.Println(string(data))
	} else {
		for i := len(contents)-1-int(*lines); i < len(contents)-1; i++ {
			fmt.Println(i+1,":",contents[i])
		}
	}
}

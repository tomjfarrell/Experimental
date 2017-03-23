// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
)

var fileflag = flag.String("file", "/var/tmp/text.rtf", "source file")
var linesflag = flag.Int("lines", 10, "lines to read, starting from end")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	flag.Parse()
	file := *fileflag
	lines := *linesflag

	fmt.Println("file flag using: ", file)

	dat, err := ioutil.ReadFile(file)
	check(err)
	contents := strings.Split(string(dat), "\n")
	if len(contents) < 11 {
		fmt.Println(string(dat))
	} else {
		for i := len(contents)-lines; i < len(contents); i++ {
			fmt.Println(contents[i])
		}
	}
}

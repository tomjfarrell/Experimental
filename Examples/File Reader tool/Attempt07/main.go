package main

import (
	"fmt"
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
func file_reader(file string, size int64, lines int) []string {
	var output []string
	var leftover string
	f, err := os.Open(file)
	check(err)
	chunk_size := 500
	data := make([]byte, chunk_size)

	lastlinecheck := make([]byte, 1)
	count, err := f.ReadAt(lastlinecheck, int64(size - 1))
	check(err)
	if string(lastlinecheck[:count]) == "\n" {
		size--
	}

	for i := int(size); i > 0 && len(output) < lines; i -= chunk_size {
		if i >= chunk_size {
			fmt.Printf("Starting read at: %v\n\n", i - chunk_size)
			count, err := f.ReadAt(data, int64(i - chunk_size))
			check(err)
			leftover = string(data[:count]) + leftover
			fmt.Printf("Read: %v\n\n", leftover)
			if strings.Contains(leftover, "\n") {
				fmt.Println("(f>c) new lines found\n")
				array := line_finder(leftover)
				leftover = array[0]
				fmt.Printf("adding to leftover: %v\n\n", array[0])
				output = append(array[1:], output...)
				fmt.Printf("output array so far (length: %v): %v\n\n", len(output), output)
			} else {
				fmt.Println("(f>c) no new lines found\n")
			}
		} else {
			remainder := make([]byte, i)
			count, err := f.Read(remainder)
			check(err)
			leftover = string(remainder[:count]) + leftover
			if strings.Contains(leftover, "\n") {
				fmt.Println("(c>f) new lines found\n")
				array := line_finder(leftover)
				leftover = array[0]
				output = append(array, output...)
				} else {
				fmt.Println("(c>f) no new lines found\n")
				output = append([]string{leftover}, output...)
			}
		}
		}
	return output
}

func line_finder(text string) []string {
	arrayout := strings.Split(text, "\n")
	fmt.Println("line_finder found these lines:")
	for i := 0; i < len(arrayout); i++{
		fmt.Printf("arrayout[%v]: %v\n\n", i, arrayout[i])
	}
	return arrayout
}

func printrequest(file string, a,z int) int64 {
	stat, err := os.Stat(file)
	check(err)
	fmt.Printf("Filename: '%s'|Size: %d bytes|Modified: %v\n", stat.Name(),stat.Size(),stat.ModTime())
	if a != 0 {
		fmt.Printf("Pulling first %v lines...\n", a)
	}
	if z != 0 {
		fmt.Printf("Pulling last %v lines...\n", z)
	}
	return stat.Size()
}

func main() {

	flag.Parse()

	size := printrequest(*file, *alines, *zlines)
	output := file_reader(*file, size, *zlines)
	for i := 0; i < len(output) && i+1 <= *zlines; i++ {
		fmt.Printf("Line #%v: %v\n\n", i+1, output[i])
	}
}
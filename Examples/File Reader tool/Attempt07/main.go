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

func back_reader(file string, size int64, lines int) []string {
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
			fmt.Printf("starting read at: %v\n\n", i - chunk_size)
			count, err := f.ReadAt(data, int64(i - chunk_size))
			check(err)
			leftover = string(data[:count]) + leftover
			fmt.Printf("read: %v\n\n", leftover)
			if strings.Contains(leftover, "\n") {
				fmt.Println("(f>c) new lines found\n")
				array := line_finder(leftover)
				leftover = array[0]
				fmt.Printf("replacing read buffer with arrayout[%v]...\n\n", 0)
				output = append(array[1:], output...)
				fmt.Printf("length of output array so far: %v\n\n", len(output))
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

func front_reader(file string, size int64, lines int) []string {
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

	for i := 0; i <= int(size) && len(output) < lines; i += chunk_size {
		if int(size) >= chunk_size {
			fmt.Printf("starting read at: %v\n\n", i)
			count, err := f.ReadAt(data, int64(i))
			check(err)
			leftover = leftover + string(data[:count])
			fmt.Printf("read buffer: %v\n\n", leftover)
			if strings.Contains(leftover, "\n") {
				fmt.Println("(f>c) new lines found\n")
				array := line_finder(leftover)
				leftover = array[len(array)-1]
				fmt.Printf("replacing read buffer with arrayout[%v]...\n\n", len(array)-1)
				output = append(output, array[:len(array)-1]...)
				fmt.Printf("length of output array so far: %v\n\n", len(output))
			} else {
				fmt.Println("(f>c) no new lines found, reading again...\n")
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

func printrequest(file string) int64 {
	stat, err := os.Stat(file)
	check(err)
	fmt.Printf("Filename: '%s'|Size: %d bytes|Modified: %v\n", stat.Name(),stat.Size(),stat.ModTime())
	return stat.Size()
}

func main() {

	flag.Parse()

	size := printrequest(*file)

	if *alines != 0 {
		fmt.Printf("Pulling first %v lines...\n", *alines)
		output := front_reader(*file, size, *alines)
		for i := 0; i < len(output) && i+1 <= *alines; i++ {
			fmt.Printf("Line #%v: %v\n\n", i + 1, output[i])
		}
	}
  if *zlines != 0 {
		fmt.Printf("Pulling last %v lines...\n", *zlines)
		output := back_reader(*file, size, *zlines)
		for i := 0; i < len(output) && i+1 <= *zlines; i++ {
			fmt.Printf("Line #%v: %v\n\n", i+1, output[i])
		}
	}
	}

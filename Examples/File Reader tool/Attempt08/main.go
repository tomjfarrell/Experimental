package main

import (
	"fmt"
	"flag"
	"strings"
	"os"
	"net/http"
	"io/ioutil"
)

var file = flag.String("file", "", "source file")
var alines =flag.Int("alines", 0, "lines to read, starting from beginning")
var zlines = flag.Int("zlines", 0, "lines to read, starting from end")
var web = flag.String("web", "", "source page")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func back_reader(size int64, lines int) []string {
	var output []string
	var leftover string

	var f *os.File
	var err error
	if *file != "" {
		f, err = os.Open(*file)
		check(err)
		defer f.Close()
	}

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

func front_reader(size int64, lines int) []string {
	var output []string
	var leftover string

	var f *os.File
	var err error
	if *file != "" {
		f, err = os.Open(*file)
		check(err)
		defer f.Close()
	}

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

func web_front_reader() []string {

	var output []string

	resp, err := http.Get(*web)
	check(err)
	defer resp.Body.Close()

	h, err := ioutil.ReadAll(resp.Body)
	check(err)

	fmt.Printf("\nWebpage: '%v'|Size: %d bytes|Modified: %v \n\n", *web, len(h), resp.Header.Get("Last-Modified"))
	fmt.Printf("Pulling first %v lines...\n\n", *alines)

	array := strings.Split(string(h), "\n")

	if len(array) >= *alines {
		output = append(array[:*alines], output...)
	}
	return output
}

func web_back_reader() []string {

	var output []string

	resp, err := http.Get(*web)
	check(err)
	defer resp.Body.Close()

	h, err := ioutil.ReadAll(resp.Body)
	check(err)

	fmt.Printf("\nWebpage: '%v'|Size: %d bytes|Modified: %v \n\n", *web, len(h), resp.Header.Get("Last-Modified"))
	fmt.Printf("Pulling last %v lines...\n\n", *zlines)

	array := strings.Split(string(h), "\n")

	if len(array) >= *zlines {
		output = append(array[len(array)-*zlines:], output...)
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

func print_file_request(file string) int64 {
	stat, err := os.Stat(file)
	check(err)
	fmt.Printf("\nFilename: '%s'|Size: %d bytes|Modified: %v\n", stat.Name(),stat.Size(),stat.ModTime())
	return stat.Size()
}

func main() {

	flag.Parse()

	if *file != "" && *web != "" {
		fmt.Printf("Cannot select web AND file.\n")
		os.Exit(3)
	}

	if *file == "" && *web == "" {
		fmt.Printf("Need to select a file OR web page.\n")
		os.Exit(3)
	}

	if *alines != 0 && *file != "" {
		size := print_file_request(*file)
		fmt.Printf("\nPulling first %v lines...\n", *alines)
		output := front_reader(size, *alines)
		for i := 0; i < len(output) && i + 1 <= *alines; i++ {
			fmt.Printf("Line #%v: %v\n\n", i + 1, output[i])
		}
	}
	if *zlines != 0 && *file != "" {
		size := print_file_request(*file)
		fmt.Printf("\nPulling last %v lines...\n", *zlines)
		output := back_reader(size, *zlines)
		for i := 0; i < len(output) && i + 1 <= *zlines; i++ {
			fmt.Printf("Line #%v: %v\n\n", i + 1, output[i])
		}
	}
	if *alines != 0 && *web != "" {
		output := web_front_reader()
		for i := 0; i < len(output) && i + 1 <= *alines; i++ {
			fmt.Printf("Line #%v: %v\n\n", i + 1, output[i])
		}
	}
	if *zlines != 0 && *web != "" {
		output := web_back_reader()
		for i := 0; i < len(output) && i + 1 <= *zlines; i++ {
			fmt.Printf("Line #%v: %v\n\n", i + 1, output[i])
		}
	}
}
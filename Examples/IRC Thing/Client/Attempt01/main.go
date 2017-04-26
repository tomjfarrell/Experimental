package main

import (
	"fmt"
	"net"
	"log"
	"os"
	"bytes"
	"bufio"
	"time"
	"flag"
	"io"
	"io/ioutil"
)

var (
	Info    *log.Logger
	Error   *log.Logger
)

var debug = flag.Bool("d", false, "set the debug modus( print informations )")
var address = flag.String("a", "127.0.0.1:9988", "set the server address")

var running bool;  // global variable if client is running

func Init(
infoHandle io.Writer,
errorHandle io.Writer) {
	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// func test(): testing for error
func test(err error, mesg string) {
	if err!=nil {
		Error.Println("CLIENT: ERROR: ", mesg)
		os.Exit(-1)
	} else {
		Info.Println("Ok: ", mesg)
	}
}

// read from connection and return true if ok
func Read(con net.Conn) string {
	buf := make([]byte, 4048)
	_, err := con.Read(buf)
	if err!=nil {
		con.Close()
		running=false
		return "Error in reading!"
	}
	str := string(buf)
	fmt.Println()
	return string(str)
}

// clientsender(): read from stdin and send it via network
func clientsender(con net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("you> ")
		input, _ := reader.ReadBytes('\n')
		if bytes.Equal(input, []byte("/quit\n")) {
			con.Write([]byte("/quit"))
			running = false
			break
		}
		Info.Println("clientsender(): send: ", string(input[0:len(input)-1]))
		con.Write(input[0:len(input)-1])
	}
}

// clientreceiver(): wait for input from network and print it out
func clientreceiver(con net.Conn) {
	for running {
		fmt.Println(Read(con))
		fmt.Print("you> ")
	}
}

func main() {
	flag.Parse();
	infolog := ioutil.Discard
	if *debug == true {
		infolog = os.Stdout
	}
	Init(infolog, os.Stderr)

	running = true
	Info.Println("main(): start ")

	// connect
	destination := *address
	Info.Println("main(): connecto to ", destination)
	con, err := net.Dial("tcp", destination)
	test(err, "dialing")
	defer con.Close()
	Info.Println("main(): connected ")

	// get the user name
	fmt.Print("Please give your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadBytes('\n')

	//cn.Write(strings.Bytes("User: "))
	con.Write(name[0:len(name)-1])

	// start receiver and sender
	Info.Println("main(): start receiver")
	go clientreceiver(con)
	Info.Println("main(): start sender")
	go clientsender(con)

	// wait for quiting (/quit). run until running is true
	for ;running; {
		time.Sleep(1*1e9)
	}
	Info.Println("main(): stoped")
}

/*
Borrowed heavily from:
http://raycompstuff.blogspot.com.au/2009/12/simpler-chat-server-and-client-in.html
https://www.goinggo.net/2013/11/using-log-package-in-go.html
*/
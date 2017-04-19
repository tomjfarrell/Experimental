package main

import (
	"fmt"
	"net"
	"os"
	"flag"
	"log"
	"io"
	"io/ioutil"
	"container/list"
)

const (
	CONN_HOST = ""
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

var running bool;  // global variable if client is running

var debug = flag.Bool("d", false, "set the debug modus( print informations )")

func Init(
traceHandle io.Writer,
infoHandle io.Writer,
warningHandle io.Writer,
errorHandle io.Writer) {
	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}



// handlingINOUT(): handle inputs from client, and send it to all other client via channels.
func handlingINOUT(IN <-chan string, lst *list.List) {
	for {
		Info.Println("handlingINOUT(): wait for input")
		input := <-IN;  // input, get from client
		// send to all client back
		Info.Println("handlingINOUT(): handling input: ", input)
		for value := range lst.Iter() {
			client := value.(ClientChat)
			Info.Println("handlingINOUT(): send to client: ", client.Name)
			client.IN<- input
		}
	}
}

// clientHandling(): get the username and create the clientsturct
// start the clientsender/receiver, add client to list.
func clientHandling(con *net.Conn, ch chan string, lst *list.List) {
	buf := make([]byte, 1024);
	con.Read(buf);
	name := string(buf);
	newclient := &ClientChat{name, make(chan string), ch, con, make(chan bool), lst};

	Info.Println("clientHandling(): for ", name);
	go clientsender(newclient);
	go clientreceiver(newclient);
	lst.PushBack(*newclient);
	ch<- name+" has joinet the chat";
}

func main() {

	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	flag.Parse()

	Info.Println("main(): start")

	// create the list of clients
	clientlist := list.New()
	in := make(chan string)
	Info.Println("main(): start handlingINOUT()")
	go handlingINOUT(in, clientlist)

	// create the connection
	netlisten, err := net.Listen("tcp", "127.0.0.1:9988")
	Error.Println(err, "main Listen")
	defer netlisten.Close()

	for {
		// wait for clients
		Info.Println("main(): wait for client ...");
		conn, err := netlisten.Accept();
		Error.Println(err, "main: Accept for client");
		go clientHandling(&conn, in, clientlist);
	}
}


/*
Borrowed heavily from:
http://raycompstuff.blogspot.com.au/2009/12/simpler-chat-server-and-client-in.html
https://www.goinggo.net/2013/11/using-log-package-in-go.html
*/

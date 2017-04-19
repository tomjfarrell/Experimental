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
// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
	////////////////////////////
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}
*/

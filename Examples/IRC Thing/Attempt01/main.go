package main

import (
	"net"
	"os"
	"flag"
	"log"
	"io"
	"io/ioutil"
	"container/list"
	"bytes"
)

var debug = flag.Bool("d", false, "set the debug modus( print informations )")

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


type ClientChat struct {
	Name string;        // name of user
	IN chan string;     // input channel for to send to user
	OUT chan string;    // input channel from user to all
	Con net.Conn;      // connection of client
	Quit chan bool;     // quit channel for all goroutines
	ListChain *list.List;    // reference to list
}

// close the connection and send quit to sender
func (c *ClientChat) Close() {
	c.Quit<-true
	c.Con.Close()
	c.deleteFromList()
}

// compare two clients: name and network connection
func (c *ClientChat) Equal(cl *ClientChat) bool {
	if bytes.Equal([]byte(c.Name), []byte(cl.Name)) {
		if c.Con == cl.Con {
			return true
		}
	}
	return false
}

// delete the client from list
func (c *ClientChat) deleteFromList() {
	for e := c.ListChain.Front(); e != nil; e = e.Next() {
		client := e.Value.(ClientChat)
		if c.Equal(&client) {
			Info.Println("deleteFromList(): ", c.Name)
			c.ListChain.Remove(e)
		}
	}
}

// read from connection and return true if ok
func (c *ClientChat) Read(buf []byte) bool{
	nr, err := c.Con.Read(buf)
	if err!=nil {
		c.Close()
		return false;
	}
	Info.Println("Read():  ", nr, " bytes")
	return true
}

// clientreceiver wait for an input from network, after geting data it send to
// handlingINOUT via a channel.
func clientreceiver(client *ClientChat) {
	buf := make([]byte, 2048)

	Info.Println("clientreceiver(): start for: ", client.Name)
	for client.Read(buf) {

		if bytes.Equal(buf, []byte("/quit")) {
			client.Close()
			break;
		}
		Info.Println("clientreceiver(): received from ",client.Name, " (", string(buf), ")")
		send := client.Name+"> "+string(buf)
		client.OUT<- send
		for i:=0; i<2048;i++ {
			buf[i]=0x00
		}
	}

	client.OUT <- client.Name+" has left chat"
	Info.Println("clientreceiver(): stop for: ", client.Name)
}

// clientsender(): get the data from handlingINOUT via channel (or quit signal from
// clientreceiver) and send it via network
func clientsender(client *ClientChat) {
	Info.Println("clientsender(): start for: ", client.Name)
	for {
		Info.Println("clientsender(): wait for input to send")
		select {
		case buf := <- client.IN:
			Info.Println("clientsender(): send to \"", client.Name, "\": ", string(buf))
			client.Con.Write([]byte(buf))
		case <-client.Quit:
			Info.Println("clientsender(): client want to quit")
			client.Con.Close()
			break;
		}
	}
	Info.Println("clientsender(): stop for: ", client.Name)
}

// handlingINOUT(): handle inputs from client, and send it to all other client via channels.
func handlingINOUT(IN <-chan string, lst *list.List) {
	for {
		Info.Println("handlingINOUT(): wait for input")
		input := <-IN;  // input, get from client
		// send to all client back
		Info.Println("handlingINOUT(): handling input: ", input)
		for e := lst.Front(); e != nil; e = e.Next() {
			client := e.Value.(ClientChat)
			Info.Println("handlingINOUT(): send to client: ", client.Name)
			client.IN<- input
		}
	}
}

// clientHandling(): get the username and create the clientsturct
// start the clientsender/receiver, add client to list.
func clientHandling(con net.Conn, ch chan string, lst *list.List) {
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

	flag.Parse()
	infolog := ioutil.Discard
	if *debug == true {
		infolog = os.Stdout
	}
	Init(infolog, os.Stderr)

	Info.Println("main(): start")

	// create the list of clients
	clientlist := list.New()
	in := make(chan string)
	Info.Println("main(): start handlingINOUT()")
	go handlingINOUT(in, clientlist)

	// create the connection
	netlisten, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	Error.Println(err, "main Listen")
	defer netlisten.Close()

	for {
		// wait for clients
		Info.Println("main(): wait for client ...")
		conn, err := netlisten.Accept()
		Error.Println(err, "main: Accept for client")
		go clientHandling(conn, in, clientlist)
	}
}


/*
Borrowed heavily from:
http://raycompstuff.blogspot.com.au/2009/12/simpler-chat-server-and-client-in.html
https://www.goinggo.net/2013/11/using-log-package-in-go.html
*/

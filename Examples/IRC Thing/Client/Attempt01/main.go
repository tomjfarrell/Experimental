package main

import (
	"fmt"
	"net"
	"log"
	"os"
	"bytes"
	"bufio"
	"strings"
	"time"
	"flag"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

var debug = flag.Bool("d", false, "set the debug modus( print informations )")

func main() {
	flag.Parse();
	running = true;
	Log("main(): start ");

	// connect
	destination := "127.0.0.1:9988";
	Log("main(): connecto to ", destination);
	cn, err := net.Dial("tcp", "", destination);
	test(err, "dialing");
	defer cn.Close();
	Log("main(): connected ");

	// get the user name
	fmt.Print("Please give you name: ");
	reader := bufio.NewReader(os.Stdin);
	name, _ := reader.ReadBytes('\n');

	//cn.Write(strings.Bytes("User: "));
	cn.Write(name[0:len(name)-1]);

	// start receiver and sender
	Log("main(): start receiver");
	go clientreceiver(&cn);
	Log("main(): start sender");
	go clientsender(&cn);

	// wait for quiting (/quit). run until running is true
	for ;running; {
		time.Sleep(1*1e9);
	}
	Log("main(): stoped");
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// only needed below for sample processing

func main() {
	fmt.Println("Launching client...")
	var host = flag.String("host", "127.0.0.1", "enter the host ip for examine the node")
	var port = flag.String("port", "8080", "enter the host port for examine the node")
	flag.Parse()
	go clientRoutine(fmt.Sprintf("%s:%s", "127.0.0.1", "8080"))
	for {
		time.Sleep(2 * time.Second)
	}

}

func clientRoutine(host string) {
	// connect to this socket
	conn, _ := net.Dial("tcp", host)
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(conn, err.Error()+"\n")
		} else {
			//fmt.Println(text)
			fmt.Fprintf(conn, text+"\n")
		}

		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

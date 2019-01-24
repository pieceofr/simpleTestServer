package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// only needed below for sample processing

func main() {

	fmt.Println("Launching server...")
	go serverRoutine("127.0.0.1", "8080")
	for {
		time.Sleep(2 * time.Second)
	}

}

//serverRoutine tcp server to recieve message continuously
func serverRoutine(host, port string) {
	ln, _ := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))

	conn, _ := ln.Accept()
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}

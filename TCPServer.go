package main

import (
	"fmt"
	"net"
	"bufio"
"strings"
)

func main() {
	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":8085")

	if err != nil{
		fmt.Print(err)
	}

	conn, err := ln.Accept()

	if err != nil{
		fmt.Print(err)
	}

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

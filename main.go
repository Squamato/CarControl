package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	tcpServer := TCPServer{localAddr: "", port: 8085}
	messages := make(chan string)
	errs := make(chan error)
	go tcpServer.start(messages, errs)

	go func() {
		for {
			fmt.Println("[Received] " + <-messages)
		}
	}()

	go func() {
		for {
			e := <-errs
			fmt.Println("[Error] " + e.Error())
		}
	}()

	for {
		s, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Print(err)
		}

		err = tcpServer.send(s)

		if err != nil {
			fmt.Print(err)
		}
	}
}

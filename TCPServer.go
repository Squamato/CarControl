package main

import (
	"bufio"
	"errors"
	"net"
	"strconv"
	"strings"
)

type TCPServer struct {
	//Local address to listen on. "" will listen on all
	localAddr  string
	port       int
	listener   net.Listener
	connection net.Conn
}

/*
	Starts listening. Can block thread due to loop.
	Messages and errors will be written to chan messages and err.
*/
func (server *TCPServer) start(messages chan string, errs chan error) {

	var err error

	server.listener, err = net.Listen("tcp", server.localAddr+":"+strconv.Itoa(server.port))

	if err != nil {
		errs <- err
	}

	server.connection, err = server.listener.Accept()

	defer server.connection.Close()
	defer server.listener.Close()

	if err != nil {
		errs <- err
	}

	for {

		message, _ := bufio.NewReader(server.connection).ReadString('\n')

		if message == "" {

			errs <- errors.New("Received empty message.")
			break
		}

		messages <- message

		newmessage := strings.ToUpper(message)

		server.connection.Write([]byte(newmessage + "\n"))
	}
}

func (server *TCPServer) send(message string) error {
	_, err := server.connection.Write([]byte(message))

	return err
}

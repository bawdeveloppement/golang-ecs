package main

import (
	"fmt"
	"net"
	"os"
)

type ForgottenKingdomServer struct {
}

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func (fks *ForgottenKingdomServer) Start() {

	listening, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
	}

	defer listening.Close()

	fmt.Println("TCP Server listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		// Listen for an incomming connection
		connection, err := listening.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(connection)
	}
}

func handleConnection(conn net.Conn) {
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

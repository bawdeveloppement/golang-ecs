package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
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
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	// Make a buffer to hold incoming data.
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Printf("-> %s", string(netData))

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			conn.Write([]byte(string("STOP\n")))
			break
		}

		result := strconv.Itoa(rand.Intn(100-0)+0) + "\n"
		conn.Write([]byte(string(result)))
	}
	conn.Close()
}

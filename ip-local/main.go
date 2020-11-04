package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read write data")
	}
}

func list() {
	listener, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatalln("unable to bind port")
	}
	log.Println("Listening on localhost:4444")

	for {
		conn, err := listener.Accept()
		log.Println("Received Connection")
		fmt.Println(conn.RemoteAddr())

		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}
}

func main() {
	fmt.Println("Starting the server...")
	list()
}

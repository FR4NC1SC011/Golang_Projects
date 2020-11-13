package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/ip2location/ip2location-go"
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
		//	strRemoteAddr := conn.RemoteAddr().String()
		addr := conn.RemoteAddr().(*net.TCPAddr).String()
		fmt.Println("Remote Address: " + addr)
		Ip2Loc(addr)

		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}
}

func Ip2Loc(RemoteAddr string) {
	db, err := ip2location.OpenDB("IP2LOCATION-LITE-DB11.BIN")
	if err != nil {
		log.Println("Error opening DataBase File")
		return
	}

	results, err := db.Get_all(RemoteAddr)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(".:INFORMATION OF THE REMOTE ADDRESS:.")
	fmt.Printf("COUNTRY SHORT %s\n", results.Country_short)
	fmt.Printf("COUNTRY LONG %s\n", results.Country_long)
	fmt.Printf("REGION %s\n", results.Region)
	fmt.Printf("CITY %s\n", results.City)

}

func main() {
	fmt.Println("Starting the server...")
	list()
}

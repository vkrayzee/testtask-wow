package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/vkrayzee/testtask-wow/services/hasher"
)

const (
	HOST = "localhost"
	PORT = 8080
	TYPE = "tcp"
)

func main() {
	// get host and port from args
	host := flag.String("host", HOST, "host")
	port := flag.Int("port", PORT, "port")

	flag.Parse()

	// connect to server
	conn, err := net.Dial(TYPE, fmt.Sprintf("%v:%v", *host, *port))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// close connection
	defer conn.Close()

	// read response from server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	challenge := string(buffer[:n])

	// print response
	// fmt.Printf("Response from server: %v\n", challenge)

	// create hasher
	h := hasher.New(nil)

	// Mint a new stamp
	stamp, _ := h.Mint(challenge)

	// send message to server
	message := stamp

	// fmt.Println("Sending message to server: " + message)
	conn.Write([]byte(message))

	n, err = conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	message = string(buffer[:n])

	// print response
	// fmt.Printf("Response from server: %v\n", message)

	fmt.Println(message)
}

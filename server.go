package main

import (
	"clat-server/connection"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Setting up logging
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// set output file
	log.SetOutput(file)

	log.Println("CLAT SERVER STARTED")

	// getting port
	var port string

	flag.StringVar(&port, "p", "8642", "Specify the port that you want to use.")
	flag.Parse()

	// Listening for connections
	server, err := net.Listen("tcp", fmt.Sprint(":", port))
	if err != nil {
		log.Fatalln("Failed to start listening. Error:\n", err.Error())
	}
	log.Println("Listening on Port:", port)

	// Main Loop
	var running = true
	for running {
		// Accept connection
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln("Cannot accept connection. Error:\n", err.Error())
			running = false
		}

		go connection.HandleConn(&conn)

	}

}

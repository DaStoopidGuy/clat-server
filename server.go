package main

import (
	"log"
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
}

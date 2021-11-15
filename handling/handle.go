package handling

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func ConnectToServer() net.Conn {
	var ip string
	var port string

	flag.StringVar(&ip, "a", "127.0.0.1", "Specify IP address. Default is localhost.")
	flag.StringVar(&port, "p", "8642", "Specify PORT. Default is 8642.")
	flag.Parse()

	// The Full Address
	var ipFull = fmt.Sprint(ip, ":", port)

	// Connecting to server
	conn, err := net.Dial("tcp", ipFull)
	if err != nil {
		log.Fatalf("Couldn't connect to the server. ERROR:\n%s", err.Error())
	}

	// returning the connection
	return conn
}

package connection

import (
	"log"
	"net"
)

func HandleConn(connection *net.Conn) {
	var conn = *connection
	log.Println("New Connection: ", conn.RemoteAddr())
}

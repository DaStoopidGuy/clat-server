package connection

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var connList []net.Conn

func HandleConn(connection *net.Conn) {
	var conn = *connection
    connList = append(connList, conn)
    
    // Create reader to read from our connection
    var reader = bufio.NewReader(conn)
    
    // log that we got a new connection
	log.Println("New Connection: ", conn.RemoteAddr())

    //Loop
    var connected = true
    for connected {
        message, err := reader.ReadString('n') 
        if err != nil {
            log.Println("Error: ", err)
            fmt.Fprintln(conn, "Error: ", err)
            connected = false
            return
        }
        //purely for the sake of debugging...😏😏
        log.Println("MessageSent: ", message)

        //send the message to all
        if message == "\n" {
        } else {
            Broadcast(message)
        }
    }
}

func Broadcast(message string) {
	for _, conn := range connList {
		fmt.Fprint(conn, message)
	}
}




package tsxcommand

import (
	"fmt"
	"net"
	"strings"
)

// TsxServer address structure
type TsxServer struct {
	Addr string
	Port int
}

// Send function to connect and send command to the TheSkyX server and javascript API interface
// returning reply as a string
func Send(tsx TsxServer, cmd string, cmdName string) string {

	// connect to this socket

	srv := fmt.Sprintf("%s:%d", tsx.Addr, tsx.Port)
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		fmt.Printf("Try tcp to server: %s\n", srv)
		fmt.Printf("error: %s\n", err)
		panic("We have no connection !")
	}
	defer conn.Close()

	// send to socket
	fmt.Printf("Get status from TheSkyX: %s\n", cmdName)
	conn.Write([]byte(cmd))

	// listen for reply
	ret := listenReply(conn)

	return string(ret)
}

func listenReply(conn net.Conn) []byte {

	// fmt.Println("Read status response")
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	status := buff[:n]

	return status[:strings.IndexByte(string(status), '|')]
}

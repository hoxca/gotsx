package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	const statusCommand = `
	  /* Java Script */
	  var Out;

		if (sky6RASCOMTele.IsConnected==0) {
		  Out = 'Mount is not connected';
		} else {
		  if (sky6RASCOMTele.IsParked()) {
			  Out = 'Mount is parked';
      } else {
			  Out = 'Mount is not parked';
			}
		} `

	const parkAndDisconnectCommand = `
	  /* Java Script */
	  var Out;
	  if (sky6RASCOMTele.IsParked()) {
      sky6RASCOMTele.Unpark();
      sky6RASCOMTele.Park();
    } else {
      sky6RASCOMTele.Park();
		}
	  Out = sky6RASCOMTele.IsConnected; `

	servAddr := "127.0.0.1:3040"

	// connect to this socket
	conn, _ := net.Dial("tcp", servAddr)
	defer conn.Close()

	// send to socket
	fmt.Println("Get status from the telescope")
	conn.Write([]byte(statusCommand))

	// listen for reply
	fmt.Println("Read status response")
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	status := buff[:n]
	ret := status[:strings.IndexByte(string(status), '|')]

	switch string(ret) {
	case "Mount is parked":
		fmt.Println("Your mount was parked")
		fmt.Println("Send disconnect to the mount")
		// send parkAndDisconnectCommand to socket
		conn.Write([]byte(parkAndDisconnectCommand))
	case "Mount is not parked":
		fmt.Println("Your mount was not parked")
		fmt.Println("Send park and disconnect to the mount")
		// send parkAndDisconnectCommand to socket
		conn.Write([]byte(parkAndDisconnectCommand))
	case "Mount is not connected":
		fmt.Println("Your mount is already disconnected")
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Printf("\n\nOutput: %s\n", status)
	}

	// listen for reply
	fmt.Println("Read status response")
	n, _ = conn.Read(buff)
	status = buff[:n]
	ret = status[:strings.IndexByte(string(status), '|')]

	if string(ret) == "0" {
		fmt.Println("Mount disconnected !")
	} else {
		fmt.Printf("Warning, Unexpected message from server: %s\n", ret)
	}
	fmt.Println("Exiting...")
}

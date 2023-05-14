// Test with netcat while running

/// Simple hello message
// echo "\x00\x01\x00\x01\x02\x01" | nc localhost 1523

package main

import (
	"fmt"
	"net"
	"os"
)

const HOST = "localhost"
const PORT = "1523"

func main() {
	listener, error := net.Listen("tcp", HOST+":"+PORT)

	if error != nil {
		fmt.Printf("Error::main::listening on: %v:%v::%v\n", HOST, PORT, error.Error())
		return
	}

	fmt.Println("main::Listening")
	for {
		cnx, e := listener.Accept()

		if e != nil {
			fmt.Printf("Error::main::%v", e.Error())
			os.Exit(1)
		}

		go Handle_request(cnx)
	}
}

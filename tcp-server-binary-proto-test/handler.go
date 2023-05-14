package main

import (
	"fmt"
	"net"
)

func handle_request(cnx net.Conn) {
	fmt.Printf("handle_request::%v\n", cnx.RemoteAddr().String())

	helloMessageBuffer := make([]byte, HELLO_MSG_LEN_IN_BYTES)

	_, e := cnx.Read(helloMessageBuffer)

	if e != nil {
		fmt.Printf("error::handle_request::%v\n", e.Error())
		cnx.Close()
		return
	}

	hello, e := parse_hello(helloMessageBuffer)
	if e != nil {
		cnx.Close()
		return
	}

	fmt.Printf("Recieved hello: %+v\n", hello)
	cnx.Write([]byte{0x00, 0x01})
	cnx.Close()
}

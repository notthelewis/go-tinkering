package main

import (
	"fmt"
	"net"
)

func Handle_request(cnx net.Conn) {
	fmt.Printf("handle_request::%v\n", cnx.RemoteAddr().String())
	// 1kb buffer
	buffer := make([]byte, 1024)

	_, e := cnx.Read(buffer)
	if e != nil {
		fmt.Printf("error::handle_request::%v\n", e.Error())
		cnx.Close()
		return
	}

	// Echo
	cnx.Write(buffer)
	cnx.Close()
}

package main

import (
    "fmt"
    "net"
    "os"
)

const HOST = "localhost"
const PORT = "1523"

func handle_request(cnx net.Conn) {
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

func main() {
    listener, error := net.Listen("tcp", HOST+":"+PORT);

    if error != nil {
        fmt.Printf("Error::main::listening on: %v:%v::%v\n", HOST, PORT, error.Error())
        return
    }

    fmt.Println("main::Listening")
    for {
        cnx, e := listener.Accept()
        if e != nil {
            fmt.Printf("Error::main::%v", e.Error());
            os.Exit(1)
        }

        handle_request(cnx)
    }
}


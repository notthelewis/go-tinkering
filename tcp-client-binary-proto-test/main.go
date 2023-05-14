package main

import (
    "fmt"
    "net"
)

const HOST = "127.0.0.1"
const PORT = "1523"

func main() {
    cnx, e := net.Dial("tcp", HOST+":"+PORT)
    if e != nil {
        fmt.Printf("ERROR::main::%v", e.Error());
    }

    helloMessage := []byte {
        0x00, 0x01, // Protocol version
        0x00, 0x01, // Software version
        0x02, // Command
        0x01, // Content-length
    }

    cnx.Write(helloMessage)
    buf := make([]byte, 2)
    cnx.Read(buf)

    fmt.Println(buf)
}

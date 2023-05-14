package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
    reader := bufio.NewReader(os.Stdin);
    fmt.Println("Type me some shi")

    line, _, e :=  reader.ReadLine()
    if e != nil {
        fmt.Println("Error!", e)
        return
    }


    fmt.Println(string(line))

}

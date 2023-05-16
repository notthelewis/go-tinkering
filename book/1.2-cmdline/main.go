// this package simply prints out its cmdline args
package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
    var s, sep string

    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }

    fmt.Println(s)
    
}

func echo2() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }

    fmt.Println(s)
}

func echo3() {
    // More concise, less allocations too.
    // A new string is not created every time 
    fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo4() {
    // All slices can be printed this way
    fmt.Println(os.Args[1:])
}

func main() {
    echo1()
    echo2()
    echo3()
}

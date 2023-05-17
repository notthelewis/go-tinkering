package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func prepend_uri_if_not_exists(filename string) string {
    if !strings.HasPrefix(filename, "http://") && !strings.HasPrefix(filename, "https://") {
        return "http://"+filename
    }

    return filename
}

func fetch_all() {
    for _, url := range os.Args[1:] {
        url = prepend_uri_if_not_exists(url)

        fmt.Printf("GET %v: ", url)
        res, e := http.Get(url)
        fmt.Printf("%v\n", res.StatusCode)

        if e != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", e)
            os.Exit(1)
        }

        f, e := os.Create("./out.txt")
        if e != nil {
            fmt.Fprintf(os.Stderr, "fetch: opening file %v\n", e)
        }

        _, err := io.Copy(f, res.Body)
        res.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Println("Finished")
    }
}

func main() {
    fetch_all()
}


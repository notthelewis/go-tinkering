// Fetchall fetches URLs in parallel and reports their times and sizes
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)


func FetchConcurrent() {
    start := time.Now()   
    ch := make(chan string)

    for _, url := range os.Args[1:] {
        go fetch(url, ch)
    }

    for range os.Args[1:] {
        fmt.Println(<-ch)
    }
    
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
    start := time.Now()

    resp, error := http.Get(url)
    if error != nil {
        ch <- fmt.Sprint(error)
        return 
    }

    nbytes, error := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if error != nil {
        ch <- fmt.Sprintf("While reading: %s: %v", url, error)
        return 
    }

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}

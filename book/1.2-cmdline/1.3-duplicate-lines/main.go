package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
    "sync"
)

func find_duplicates1() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)


    for input.Scan() {
        if input.Text() == "exit" {
            break 
        }

        counts[input.Text()]++
    }

    fmt.Println(counts)

    // Ignoring potential errors from input.Err()
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

// This one can read from any file descriptor
func find_duplicates2() {
    counts := make(map[string]int)
    files := os.Args[1:]

    if len(files) == 0 {
        count_lines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, e := os.Open(arg)
            if e != nil {
                fmt.Fprintf(os.Stderr, "duplicates2: %v\n", e)
            }
            count_lines(f, counts)
            f.Close()
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func count_lines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
}

func find_duplicates_with_file() {
    counts := make(map[string]map[string]int)
    files := os.Args[1:]

    // I had to google to find out how to make coroutines wait... 
    var wg sync.WaitGroup

    for _, file := range files {
        f, err := os.Open(file)
        if err != nil {
            fmt.Fprintf(os.Stderr, "find_duplicates_with_file: %v\n", err)
            continue
        }

        wg.Add(1)
        go count_lines_with_file(f, counts, &wg)
    }

    wg.Wait() 
    for file, counter := range counts {
        for line, times_occured := range counter {
            if times_occured > 1 {
                fmt.Printf("%v had a line: %v which had %v occurances\n", file, line, times_occured)
            }
        }
    }
}

func count_lines_with_file(f *os.File, counts map[string]map[string]int, wg *sync.WaitGroup) {
    defer wg.Done()
    counts[f.Name()] = make(map[string]int)

    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[f.Name()][input.Text()]++
    }
}

func find_duplicates3() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)

        if err != nil {
            fmt.Fprintf(os.Stderr, "duplicates3: %\n", err)
            continue
        }

        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
    }

    for line, n := range counts { 
        if n > 1 { 
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}


func main() {
    // find_duplicates1()
    // find_duplicates2()
    //find_duplicates3()
    find_duplicates_with_file()
}

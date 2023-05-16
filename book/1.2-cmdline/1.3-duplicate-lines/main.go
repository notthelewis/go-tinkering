package main 

import (
    "bufio"
    "fmt"
    "os"
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

func main() {
    // find_duplicates1()
    find_duplicates2()
}

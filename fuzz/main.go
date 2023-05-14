package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("Input is not valid UTF-8");
    }
    r := []rune(s)

    for i, j := 0, len(r)-1; i < len(r) / 2; i,j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }

    return string(r), nil
}

func main() {
    input := "The quick brown fox jumped over the lazy dog"

    rev, e:= Reverse(input)
    if e != nil {
         fmt.Printf("Error with rev: %q", e)
         return;
    } 

    doubleRev, e := Reverse(rev)

    if e != nil {
        fmt.Printf("Error with double rev: %q", e)
        return
    }

    fmt.Printf("Original: %q\nReversed: %q\nDoubleReversed: %q\n",
        input,
        rev,
        doubleRev,
    )
}

package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
    testCases := []struct {
        in, want string
    }{
        {"Hello world", "dlrow olleH"},
        {" ", " "},
        {"!1234", "4321!"},
    }

    for _, tc := range testCases {
        rev, _ := Reverse(tc.in)
        if rev != tc.want {
            t.Errorf("Reverse: %q, want %q", rev, tc.want)
        }
    }
}

// To use fuzzing:
//      -> go test -fuzz=Fuzz -fuzztime=10s
func FuzzReverse(f *testing.F) {
    testCases := []string{"Hello, world", " ", "!12345", "test", "asdasd", "ki91jasd"}

    for _, tc := range testCases {
        f.Add(tc)
    }

    f.Fuzz(func(t *testing.T, orig string) {
        rev, _ := Reverse(orig)
        doubleRev, _:= Reverse(rev)

        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }

        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string: %q", rev)
        }
    })
}

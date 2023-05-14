package main

import "fmt"

func swap(a, b *uint8) {
    *a ^= *b;
    *b ^= *a;
    *a ^= *b;
}

func bubble_sort(m []uint8) {
    swapped := true
    
    // Works like a do-while
    for swapped == true {
        swapped = false

        for i, j := 0, 1; i < len(m)-1; i, j = i+1, j+1 {
            if m[i] > m[j] {
                swap(&m[i], &m[j])
                swapped = true
            }
        }
    }
}

func main() {
    var toSort = []uint8{
        10,
        3,
        13,
        100,
        4,
        10,
        7,
        5,
        255,
        9,
    }

    fmt.Println(toSort)
    bubble_sort(toSort)
    fmt.Println(toSort)
}

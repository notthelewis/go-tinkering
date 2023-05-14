package main

import (
	"fmt"
	"math"
)



func binary_search(arr []uint, toSearch uint) int {
    if len(arr) <= 1 {
        return -1 
    }

    halfLen := uint(math.Floor(float64(len(arr))/2))
    
    if arr[halfLen] == toSearch {
        return int(halfLen)
    } else if  arr[halfLen] < toSearch {
        return binary_search(arr[halfLen : len(arr)-1], toSearch)
    } else {
        return binary_search(arr[0: halfLen], toSearch)
    }

}

func main() {
    toSearch := 10
    arr := []uint{
        100,
        200,
        30,
        40,
        50,
        10,
        91,
        244,
        31,
        44,
        3,
    }
    
    fmt.Println(binary_search(arr, uint(toSearch)))
}

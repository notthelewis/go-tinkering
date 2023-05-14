package main

import "fmt"

type Number64 interface {
    int64 | float64
}

func sumInts(ints map[string]int64) int64 {
	var sum int64

	for _, int := range ints {
		sum += int
	}

	return sum
}

func sumFloats(floats map[string]float64) float64 {
	var sum float64

	for _, float := range floats {
		sum += float
	}

	return sum
}

func genericSum[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func genericPrintAll[K comparable, V any](m map[K]V) {
	for i, v := range m {
		fmt.Println(i, v)
	}
}

func sum64s[K comparable, V Number64](m map[K]V) V {
    var sum V

    for _, v := range m {
        sum += v
    }

    return sum
}

func main() {
	ints := map[string]int64{
		"first":  35,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.9,
		"second": 12.3,
	}

	fmt.Printf("Non-generic sums: %v & %v\n",
		sumInts(ints),
		sumFloats(floats),
	)

	fmt.Printf("genericSums: %v & %v \n",
		genericSum(floats),
		genericSum(ints),
	)

    genericPrintAll(ints)
    genericPrintAll(floats)

	fmt.Printf("genericSums: %v & %v \n",
		sum64s(floats),
		sum64s(ints),
	)
}

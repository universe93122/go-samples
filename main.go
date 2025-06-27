package main

import "fmt"

type Number interface {
	int | float64
}

func SumInts(m map[string]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumFloats(m map[string]float64) float64 {
	sum := 0.0
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	sum := V(0)
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	// Example usage
	// ints := map[string]int{"a": 1, "b": 2, "c": 3}
	// floats := map[string]float64{"x": 1.1, "y": 2.2, "z": 3.3}

	// fmt.Println("Sum of ints:", SumInts(ints))
	// fmt.Println("Sum of floats:", SumFloats(floats))

	// fmt.Printf("Generic Sums: %v and %v\n",
	// 	SumIntsOrFloats(ints),
	// 	SumIntsOrFloats(floats))

	// Reverse String

	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

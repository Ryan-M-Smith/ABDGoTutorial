//
// Filename: exercise04.go | Alpha Beta Debug Go Tutorial
// Description: Array and slice syntax practice
// Copyright (c) 2025 Ryan Smith
//

package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare an array and a slice to work with
	array := [4]int {0, 1, 2, 3}
	slice := []float32 {0.4, 0.1, 0.9}

	// Print info about the array
	fmt.Printf("Array: %v\n", array)
	fmt.Printf("Length: %d\n", len(array))
	fmt.Printf("Capacity: %d\n", cap(array))

	// Print info about the slice
	fmt.Printf("\nSlice: %v\n", slice)
	fmt.Printf("Length: %d\n", len(slice))
	fmt.Printf("Capacity: %d\n", cap(slice))

	// Indexing and slicing
	fmt.Printf("array[2] = %d", array[2])
	fmt.Printf("slice[:2] = %v", slice[:2])
	
	// Append to the slice, then check again - the capacity is higher than the length
	slice = append(slice, math.Pi)

	fmt.Printf("\nSlice: %v\n", slice)
	fmt.Printf("Length: %d\n", len(slice))
	fmt.Printf("Capacity: %d\n", cap(slice))

	// Slice an array to get a new slice. The data type of slicedArr should be []int
	slicedArr := array[1:3]
	fmt.Printf("\n%v (type = %T)\n", slicedArr, slicedArr)

	// Create a deep copy of a slice
	cpy := make([]float32, len(slice))
	copy(cpy, slice)

	//
	// Modify the copy - the original slice is not changed
	//

	for i := range slice {
		cpy[i] *= 2
	}

	fmt.Printf("\nOriginal: %v\n", slice)
	fmt.Printf("Copy: %v\n", cpy)
}
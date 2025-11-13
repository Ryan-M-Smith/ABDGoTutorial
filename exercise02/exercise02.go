//
// Filename: exercise02.go | Alpha Beta Debug Go Tutorial
// Description: Variables, data types, and I/O practice
// Copyright (c) 2025 Ryan Smith
//

package main

import (
	"fmt"
)

func main() {
	var a int = 5
	fmt.Println(a)

	greeting := "Hello"
	var name string
	fmt.Print("\nWhat is your name? ")
	fmt.Scanf("%s", &name)
	fmt.Printf("%s, %s!\n", greeting, name)

	salary := 120_000.00
	fmt.Printf("\nI make $%d/year as a software engineeer!", int(salary))


	i := 0

	for {
		fmt.Println(i)
		i++

		if i >= 10 {
			break
		}
	}




}
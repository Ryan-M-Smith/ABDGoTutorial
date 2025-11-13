//
// Filename: exercise05.go | Alpha Beta Debug Go Tutorial
// Description: Inventory manager (function practice)
// Copyright (c) 2025 Ryan Smith
//

package main

import (
	"fmt"
)

// Add an item to the inventory
func add(inventory []string, item string) []string {
	return append(inventory, item)
}

// List all items in the inventory
func list(inventory []string) {
	if len(inventory) == 0 {
		fmt.Println("Your inventory is currently empty.")
		return
	}

	fmt.Println("Inventory:")
	for i, item := range inventory {
		fmt.Printf("%d: %s\n", i + 1, item)
	}
}

// Remove an item from the inventory by index
func remove(inventory []string, index int) []string {
	if index >= 0 && index < len(inventory) {
		inventory = append(
			inventory[:index],
			inventory[index+1:]...
		)
	} else {
		fmt.Println("Invalid index")
	}

	return inventory
}

func main() {
	var inventory []string
	var input string 

	fmt.Println("Inventory Manager")
	
	for input != "quit" {
		fmt.Print("Enter command (add, list, remove, quit): ")
		fmt.Scanln(&input)

		// Switch on the user input to determine what to do
		switch input {
			case "add":
				var item string
				fmt.Print("Item name: ")
				fmt.Scanln(&item)
				inventory = add(inventory, item)

			case "list":
				list(inventory)

			case "remove":
				var index int
				fmt.Printf("Index to remove (1-%d): ", len(inventory))
				fmt.Scanf("%d", &index)

				// Convert to zero-based index
				inventory = remove(inventory, index - 1)

			case "quit":
				fmt.Println("Goodbye!")

			default:
				fmt.Println("Unknown command")
		}
	}
}
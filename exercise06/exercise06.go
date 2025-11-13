//
// Filename: exercise05.go | Alpha Beta Debug Go Tutorial
// Description: Inventory manager (function practice)
// Copyright (c) 2025 Ryan Smith
//

package main

import (
	"fmt"
	"math/rand"
)

// A person
type Person struct {
	firstName 	string
	lastName  	string
	age 		uint
	job 		string
}

// Print info about a person
func (p Person) Print() {
	fmt.Printf(
		"%s %s is %d years old and is a %s\n",
		p.firstName,
		p.lastName,
		p.age,
		p.job,
	)
}

// Generate random people and send them to a channel
func makePeople(people chan<- Person, numPeople int) {
	firstNames 	:= [5]string {"Liam", "Olivia", "Noah", "Ava", "Ethan"}
	lastNames 	:= [5]string {"Smith", "Johnson", "Brown", "Taylor", "Anderson"}
	ages 		:= [5]uint 	 {24, 31, 19, 28, 35}
	jobs 		:= [5]string {"Engineer", "Teacher", "Designer", "Chef", "Doctor"}
	
	for range numPeople {
		// Get random attributes
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastNames := lastNames[rand.Intn(len(lastNames))]
		age 	  := ages[rand.Intn(len(ages))]
		job 	  := jobs[rand.Intn(len(jobs))]

		// Create person and send it to the channel
		person := Person{firstName, lastNames, age, job}
		people <- person
	}

	close(people) // Close the channel when done
}

func main() {
	people := make(chan Person)
	numPeople := 10

	// Execute a goroutine
	go makePeople(people, numPeople)

	// Print each person as they are sent to the channel
	for person := range people {
		person.Print()
	}
}
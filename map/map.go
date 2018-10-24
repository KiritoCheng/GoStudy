package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["001"] = PersonInfo{"001", "Tom", "Room 203,..."}
	personDB["003"] = PersonInfo{"003", "Jack", "Room 101,..."}

	person, ok := personDB["001"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID 001.")
	} else {
		fmt.Println("Did not find person with ID 001.")
	}
}

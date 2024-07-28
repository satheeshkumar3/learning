package main

import "fmt"

func set() {
	set := make(map[string]struct{})
	set["apple"] = struct{}{}
	set["banana"] = struct{}{}
	set["cherry"] = struct{}{}

	// Check if an item is in the set
	if _, exists := set["banana"]; exists {
		fmt.Println("banana is in the set")
	}

	// Iterating over the set
	for item := range set {
		fmt.Println(item)
	}
}

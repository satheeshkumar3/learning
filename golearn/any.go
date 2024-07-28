package main

import "fmt"

func anyGo() {
	var i any = "Hello, Go!"

	// Type assertion
	s, ok := i.(string)
	if ok {
		fmt.Println(s) // Output: Hello, Go!
	} else {
		fmt.Println("Type assertion failed")
	}

	// Direct type assertion (without checking ok)
	// Will panic if the assertion fails
	s = i.(string)
	fmt.Println(s) // Output: Hello, Go!
}

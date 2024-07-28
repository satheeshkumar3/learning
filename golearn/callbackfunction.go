package learning

// Function that takes a callback function as an argument
func compute(value int, callback func(int) int) int {
	return callback(value)
}

// Example callback functions
func square(n int) int {
	return n * n
}

func double(n int) int {
	return n * 2
}

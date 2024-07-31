package golearn

// Function that takes a callback function as an argument
func Compute(value int, callback func(int) int) int {
	return callback(value)
}

// Example callback functions
func Square(n int) int {
	return n * n
}

func Double(n int) int {
	return n * 2
}

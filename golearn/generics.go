package golearn

import (
	"fmt"
)

type Numbers interface {
	int | int64 | float64
}

func Generic[T Numbers](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func SampleInput() {
	numbers := []int{10, 20, 30}
	fmt.Println(Generic(numbers))
}

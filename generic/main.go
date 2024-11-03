package main

import (
	"fmt"
	"strings"
)

func Contains[T comparable](list []T, item T) bool {
	for _, element := range list {
		if element == item {
			return true
		}
	}
	return false
}

func KeyExists[K comparable, V any](m map[K]V, key K) bool {
	_, exists := m[key]
	return exists
}

func Apply[T any](values []T, operation func(T) T) []T {
	results := make([]T, len(values))
	for i, v := range values {
		results[i] = operation(v)
	}
	return results
}

func Apply2[T any](values []T, operation func(...T) T) T {
	return operation(values...)
}

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Contains(numbers, 3)) // return true
	fmt.Println(Contains(numbers, 6)) // return false

	words := []string{"apple", "banana", "cherry"}
	fmt.Println(Contains(words, "banana")) // return true
	fmt.Println(Contains(words, "grape"))  // return false

	intMap := map[string]any{
		"age":  "30",
		"year": 2023,
	}

	fmt.Println(KeyExists(intMap, "age"))   // return true
	fmt.Println(KeyExists(intMap, "month")) // return false

	double := func(x int) int {
		return x * 2
	}
	fmt.Println(Apply(numbers, double)) // return  [2, 4, 6, 8, 10]

	concat := func(parts ...string) string {
		return strings.Join(parts, " ")
	}

	fmt.Println(Apply2(words, concat)) // return apple banana cherry
}

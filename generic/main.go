package main

import "fmt"

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
}

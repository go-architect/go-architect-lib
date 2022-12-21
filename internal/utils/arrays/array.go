// Package arrays provides utility functions for arrays
package arrays

import "fmt"

// RemoveDuplicatedStrings receives an slice/array and removes the duplicated values
func RemoveDuplicatedStrings(array []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range array {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// IndexOf returns the index of the desired value in the slice/array.
// If the value is not present in the collection, then IndexOf will return -1
func IndexOf(array []string, value string) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

// Dequeue supports queue operations, it returns the first value in a slice and removes it from the slice
func Dequeue(queue []string) (string, error) {
	if len(queue) == 0 {
		return "", fmt.Errorf("queue is empty")
	}
	element := queue[0]
	queue = queue[1:]
	return element, nil
}

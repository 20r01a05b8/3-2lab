package main

import "fmt"

// function to remove duplicate values
func removeDuplicates(s []string) []string {
	bucket := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := bucket[str]; !ok {
			bucket[str] = true
			result = append(result, str)
		}
	}
	return result
}
func main() {
	// creating an array of strings
	array := []string{"hai", "cse", "batch", "2020", "cse", "hai", "section"}
	fmt.Println("The given array of string is:", array)
	fmt.Println()

	// calling the function
	result := removeDuplicates(array)
	fmt.Println("The array obtained after removing the duplicate entries is:", result)
}

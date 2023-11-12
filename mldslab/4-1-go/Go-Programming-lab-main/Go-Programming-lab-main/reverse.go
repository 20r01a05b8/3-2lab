package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Interger Reverse Sort")
	num := []int{50, 90, 30, 10, 50}
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)
	fmt.Println()
	fmt.Println("String Reverse Sort")
	text := []string{"Japan", "UK", "Germany", "Australia", "INDIA"}
	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)
}

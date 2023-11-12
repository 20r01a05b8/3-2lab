package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Germany", "G"))
	fmt.Println(strings.ContainsAny("Germany", "g"))
	fmt.Println(strings.Contains("India", "In"))
	fmt.Println(strings.Contains("Germany", "ger"))
	fmt.Println(strings.Contains("Japan", "Jap"))
	fmt.Println(strings.Count("malayalam", "m"))
	fmt.Println(strings.Count("malayalam", "o"))
	fmt.Println(strings.EqualFold("Cat", "cAt"))
	fmt.Println(strings.EqualFold("India", "Indiana"))
}

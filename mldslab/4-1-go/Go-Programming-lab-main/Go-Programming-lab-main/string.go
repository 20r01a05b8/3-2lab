package main

import "fmt"

func main() {
	fmt.Print("Enter First String: ")
	var first string
	fmt.Scanln(&first)
	fmt.Print("Enter Second String: ")
	var second string
	fmt.Scanln(&second)
	fmt.Print(first + second)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "./tree"
)

func main() {
	// Get data
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Data: ")
	scanner.Scan()
	input := scanner.Text()
	words := strings.Split(input, " ")

	// Create tree
	root := CreateNode()

	for _, word := range words {
		root.AddWord(word)
	}

	fmt.Print("Search: ")
	scanner.Scan()
	input = scanner.Text()

	suggestions := root.GetSuggestions(input)

	fmt.Println(suggestions)
}

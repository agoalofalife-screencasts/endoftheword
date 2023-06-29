package main

import "fmt"

func main() {
	s := findMatchLettes("golang", "language")
	fmt.Println(s)
}

func findMatchLettes(firstWord string, secondWord string) int {
	index := len(firstWord) - 1

	for index >= 0 {
		if firstWord[index] == secondWord[0] {
			return index
		}
		index--
	}
	return -1
}

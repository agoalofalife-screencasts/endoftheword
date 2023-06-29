package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
var history string

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func main() {
	fmt.Println()
	fmt.Println(Blue + "Welcome to the Game 'End of the Word' by English ❤️" + Reset)
	fmt.Println()

	fmt.Println(Cyan + "The rules are too easy. You have to continue the next English word which end of previous one." + Reset)
	fmt.Println(Cyan + "For example:" + Reset)
	fmt.Println(Cyan + "I write:" + Reset)

	simulatePrinting(Red + "golang..." + Reset)
	fmt.Println(Cyan + "and your word has to start of end previous word" + Reset)
	fmt.Println(Cyan + "for example" + Reset)
	simulatePrinting(Red + "language..." + Reset)
	fmt.Println(Cyan + "It means like..." + Reset)
	simulatePrinting("go" + Red + "|" + Reset + "language")

	fmt.Println(Cyan + "Ок? easy peasy!" + Reset)
	fmt.Println(Cyan + "Let's get started! 🎉🎉🎉" + Reset)
	fmt.Println()

	for {
		var word string
		fmt.Print(Yellow + "Enter your word: " + Reset)
		fmt.Scan(&word)

		if word != "" {
			addToHistory(word)
			fmt.Println()
			fmt.Println(history)
			fmt.Println()
		}
	}
}

func addToHistory(word string) {
	countFailAttempt := 2

	isFirstWord := history == ""
	if isFirstWord {
		history = word
		return
	}
	for countFailAttempt != 0 {
		index := isCorrect(word, countFailAttempt)

		isSuccessAttempt := index != -1
		if isSuccessAttempt {
			addLastWordToHistoryString(index, word)
			countFailAttempt = 0
		} else {
			countFailAttempt--
		}
	}
}

func addLastWordToHistoryString(startPosition int, newWord string) {
	history = history[0:startPosition] + Red + "|" + Reset + newWord
}

func isCorrect(word string, orderMatch int) int {
	firstMatchedLetter := findMatchBy(orderMatch, history, word)

	for i := firstMatchedLetter + 1; i <= len(history)-1; i++ {
		if history[i] != word[i-firstMatchedLetter] {
			return -1
		}
	}
	return firstMatchedLetter
}

func findMatchBy(counterMatch int, firstWord string, secondWord string) int {
	index := len(firstWord) - 1

	for index >= 0 {
		if firstWord[index] == secondWord[0] {
			counterMatch--
			if counterMatch == 0 {
				return index
			}
		}
		index--
	}
	return -1
}

func simulatePrinting(word string) {
	const min int = 100
	const max int = 300

	for _, letter := range word {
		fmt.Printf("%c", letter)
		coefficient := rand.Intn(max-min) + min
		duration := time.Millisecond * time.Duration(coefficient)
		time.Sleep(duration)
	}

	fmt.Println("")
}

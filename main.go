package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	// Generate a new word for hangman
	resp, err := http.Get("https://random-word-api.vercel.app/api?words=1")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	word := removeNonAlphabetic(string(body))
	obscuredWord := obscureWord(word)
	guess := ""
	fmt.Printf("Word: %s. Obscured: %s\n", word, obscuredWord)

	fmt.Println("Welcome to Hangman!")
	fmt.Println("You have 6 lives to guess the word.")
	lives := 6

	for lives > 0 {
		fmt.Printf("Lives: %d. Word: %s\n", lives, obscuredWord)
		fmt.Println("Enter a letter: ")
		fmt.Scanln(&guess)

		if strings.Contains(word, guess) {
			obscuredWord = revealLetters(word, obscuredWord, guess)
		} else {
			lives--
		}
	}

}

func revealLetters(word, obscuredWord string, guess string) string {
	runeSlice := []rune(obscuredWord)
	for i, letter := range word {
		if string(letter) == guess {
			runeSlice[i] = letter
		}
	}
	return string(runeSlice)
}

// Replaces every letter with underscores
func obscureWord(word string) string {
	obscuredWord := ""
	for range word {
		obscuredWord += "_"
	}
	return obscuredWord
}

// Filters out any uppercase or non-alphabetic characters
func removeNonAlphabetic(word string) string {
	re := regexp.MustCompile("[^a-z]")
	return re.ReplaceAllString(word, "")
}

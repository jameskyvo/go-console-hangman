package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {
	// Generate a new word for hangman
	resp, err := http.Get("https://random-word-api.vercel.app/api?words=1")
	if err != nil {
		log.Fatal(err)
	}

	// Print the word
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	word := removeNonAlphabetic(string(body))
	obscuredWord := obscureWord(word)

	fmt.Printf("Word: %s. Obscured: %s\n", word, obscuredWord)
}

func obscureWord(word string) string {
	obscuredWord := ""
	for range word {
		obscuredWord += "_"
	}
	return obscuredWord
}

func removeNonAlphabetic(word string) string {
	// Filter out any uppercase or non-alphabetic characters
	re := regexp.MustCompile("[^a-z]")
	return re.ReplaceAllString(word, "")
}

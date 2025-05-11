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
	willRepeat := true
	for willRepeat {
		var (
			word         = generateWord()
			obscuredWord = obscureWord(word)
			userInput    = ""
			lives        = 6
		)

		fmt.Println("Welcome to Hangman!")
		fmt.Printf("You have %d lives to guess the word. \n", lives)

		for lives >= 0 {
			if lives == 0 {
				loseGame(word, &willRepeat)
				break
			} else if obscuredWord == word {
				winGame(word, &willRepeat)
				break
			}

			fmt.Printf("Lives: %d. Word: %s\n", lives, obscuredWord)
			fmt.Println("Enter a letter: ")
			fmt.Scanln(&userInput)

			if strings.Contains(word, userInput) {
				obscuredWord = revealLetters(word, obscuredWord, userInput)
			} else {
				lives--
			}
		}
	}

}

// Shows victory dialogue and sets willRepeat accordingly
func winGame(word string, willRepeat *bool) {
	var userInput string
	fmt.Println("Congratulations! You've guessed the word: ", word)
	fmt.Println("Would you like to play again? (y/n)")
	fmt.Scanln(&userInput)
	RepeatGame(userInput, willRepeat)
}

// Shows losing dialogue and sets willRepeat accordingly
func loseGame(word string, willRepeat *bool) {
	var userInput string
	fmt.Println("Game over! The word was: ", word)
	fmt.Println("Would you like to play again? (y/n)")
	fmt.Scanln(&userInput)
	RepeatGame(userInput, willRepeat)
}

// Restarts the game if the user inputs "y"
func RepeatGame(userInput string, willRepeat *bool) {
	filteredInput := strings.ToLower(string(userInput[0]))
	if filteredInput != "y" {
		*willRepeat = false
	}
}

// Returns a random formatted word from the API
func generateWord() string {
	resp, err := http.Get("https://random-word-api.vercel.app/api?words=1")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return removeNonAlphabetic(string(body))
}

// Reveals letters in obscuredWord if the guess is correct
func revealLetters(word, obscuredWord string, guess string) string {
	runeSlice := []rune(obscuredWord)
	if guess == word {
		return word
	}

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

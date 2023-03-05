package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"

	"github.com/fatih/color"
)

const maxAttempts = 6

type WordData []string

type WordApiResponse []struct {
	Word          string `json:"word"`
	Definition    string `json:"definition"`
	Pronunciation string `json:"pronunciation"`
}

type GuessRecordList []struct {
	Guess  string
	Result []string
}

func askGuess() string {

	var guess string

	for {
		fmt.Scan(&guess)
		guess = strings.ToUpper(guess)

		// Check if the guess is 5 letters long
		if len(guess) != 5 {
			fmt.Println("⚠️  La palabra debe tener 5 letras.")
			continue
		}
		break
	}

	return guess
}

// Check if the guess is correct and return a slice of the status of each word.
func checkGuess(guess string, word string) (status []string) {

	// Check each character of guess and compare with word
	for i := 0; i < len(guess); i++ {
		if guess[i] == word[i] {
			status = append(status, "correct")
		} else if strings.ContainsAny(word, string(guess[i])) {
			status = append(status, "present")
		} else {
			status = append(status, "absent")
		}
	}

	return

}

// Prints the status of each word.
func printCheckedGuess(guess string, guessResult []string) {

	stringTemplate := " %s "

	for i, status := range guessResult {

		character := string(guess[i])

		if status == "correct" {
			color.New(color.FgGreen, color.Bold).Printf(stringTemplate, character)
		} else if status == "present" {
			color.New(color.FgGreen, color.Bold).Printf(stringTemplate, character)
		} else {
			color.New(color.FgRed, color.Bold).Printf(stringTemplate, character)
		}

	}

	fmt.Println("")

}

func main() {

	// Read json with word data. Original: https://gist.github.com/NiciusB/860b1e5b73f95fbb2c49
	jsonFile, err := os.Open("wordData.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var wordData WordData
	json.Unmarshal(byteValue, &wordData)

	// Get random word
	word := wordData[rand.Intn(len(wordData))]
	word = strings.ToUpper(word)

	// fmt.Println("DEBUG:", word) //! DEBUG
	fmt.Printf("Adivina la palabra de 5 letras en %d intentos.\n", maxAttempts)

	attempts := 1
	var guessRecordList GuessRecordList

	for {

		fmt.Printf("👉 Intento %d/%d: ", attempts, maxAttempts)
		guess := askGuess()

		guessResult := checkGuess(guess, word)

		guessRecordList = append(guessRecordList, struct {
			Guess  string
			Result []string
		}{
			Guess:  guess,
			Result: guessResult,
		})

		printCheckedGuess(guess, guessResult)

		fmt.Println("")

		if guess == word {
			fmt.Println("✅ ¡Correcto!")
			fmt.Println("🏆 Has acertado en ", attempts, " intentos.")
			break
		}

		if attempts == maxAttempts {
			fmt.Println("💀 Has perdido")
			fmt.Println("La palabra era: ", word)
			break
		}

		attempts++
	}

	// Print results
	fmt.Println("\nTus resultados:")

	for _, guessResult := range guessRecordList {
		printCheckedGuess(guessResult.Guess, guessResult.Result)
	}

	// Define a channel to receive signals
	c := make(chan os.Signal, 1)

	// Notify any interrupt signal to the channel c
	signal.Notify(c, os.Interrupt)

	// Wait until the signal is received
	fmt.Println("\nPresiona Ctrl+C para salir")

	// Block until a signal is received.
	<-c

}

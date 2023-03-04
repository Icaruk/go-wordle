package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type WordData []string

type WordApiResponse []struct {
	Word          string `json:"word"`
	Definition    string `json:"definition"`
	Pronunciation string `json:"pronunciation"`
}

func askGuess() string {

	var guess string

	for {
		fmt.Print("👉 ")
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

	// "correct" → 🟩
	// "present" → 🟨
	// "absent" → ⚫

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	// Iterate over guessResult using range
	for i, status := range guessResult {

		character := string(guess[i])

		if status == "correct" {
			fmt.Printf("%s ", green(character))
		} else if status == "present" {
			fmt.Printf("%s ", yellow(character))
		} else {
			fmt.Printf("%s ", red(character))
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

	// fmt.Println("DEBUG:", word) //TODO: remove me
	fmt.Println("Adivina la palabra (5 letras): ")

	attempts := 1

	for {

		// fmt.Println("Intento número ", attempts)

		guess := askGuess()
		guessResult := checkGuess(guess, word)
		printCheckedGuess(guess, guessResult)

		fmt.Println("")

		if guess == word {
			fmt.Println("✅ ¡Correcto!")
			fmt.Println("🏆 Has acertado en ", attempts, " intentos.")

			fmt.Println("Cerrando en 5 segundos...")
			time.Sleep(5 * time.Second)
			break
		}

		attempts++
	}

}

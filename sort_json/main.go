package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

const inputFile = "wordData.json"
const outputFile = "wordDataSorted.json"

type WordData []string

func main() {

	// Read the file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	// Unmarshal the file
	var wordData WordData
	json.Unmarshal(byteValue, &wordData)

	// Sort the words
	sort.Strings(wordData)

	// Write the sorted file
	file, err = os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ = json.Marshal(wordData)

	// Write the sorted file
	_, err = file.Write(byteValue)
	if err != nil {
		log.Fatal(err)
	}

}

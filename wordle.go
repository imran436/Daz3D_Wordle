package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
)

const wordsLink = "https://raw.githubusercontent.com/tabatkins/wordle-list/main/words"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorReset = "\033[0m"

func startGame(wordList []string) {
	fmt.Printf("Welcome to Wordle!\n")
	var word string = getWord(wordList)
	var won bool = false
	for r := 0; r < 6; r++ {
		fmt.Println("Round ", r+1, " of 6")
		fmt.Print(">")
		var guess string
		fmt.Scanln(&guess)
		if !validWord(wordList, guess) {
			fmt.Println("That guess is invalid")
			fmt.Println("Please enter a valid 5 letter word")
			r -= 1
			continue
		}
		if guess == word {
			won = true
			winGame(word)
			break
		} else {
			resolveGuess(guess, word)
		}
	}
	if !won {
		fmt.Printf("better luck next time! \n The correct answer was \n" + word)
	}
}
func resolveGuess(guess string, word string) {
	var outPut string
	for i := 0; i < 5; i++ {
		switch {
		case guess[i] == word[i]:
			outPut += colorGreen + string(guess[i]) + colorReset
		case strings.Contains(word, string(guess[i])):
			outPut += colorYellow + string(guess[i]) + colorReset
		default:
			outPut += string(guess[i])
		}
	}
	fmt.Printf(outPut + "\n")
}
func main() {
	// add while playing loop and user input prompt after winning
	resp, err := http.Get(wordsLink)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	wordList := strings.Split(string(body), "\n")
	sort.Sort(sort.StringSlice(wordList))
	startGame(wordList)

}

func getWord(wordList []string) string {
	rand.Seed(time.Now().UnixNano())
	word := wordList[rand.Intn(len(wordList)-1)+1]
	return word
}

func winGame(answer string) {
	fmt.Println("\033[42m\033[30m" + answer + "\033[0m")
	fmt.Println("You Win!")
}

func validWord(wordList []string, guess string) bool {
	if len(guess) != 5 {
		return false
	}
	index := sort.SearchStrings(wordList, guess)
	return wordList[index] == guess
}

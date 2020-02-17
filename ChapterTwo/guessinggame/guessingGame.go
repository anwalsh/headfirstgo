package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	const limit = 101
	game(limit)
}

func genTarget(upper int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(upper)
}

func getGuess() int64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Guess a whole number between 1-100:")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}

	input = strings.TrimSpace(input)
	guess, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		fmt.Fprintln(os.Stderr, "parsing an int: ", err)
	}
	return guess
}

func compare(guess int64, target int64) bool {
	if guess > target {
		fmt.Println("Oops. Your guess was HIGH.")
	} else if guess < target {
		fmt.Println("Oops. Your guess was LOW.")
	}
	return guess == target
}

func game(limit int64) {
	target := genTarget(int64(limit))
	guesses := 0

	for guesses != 10 {
		result := compare(getGuess(), target)
		if result {
			fmt.Println("Good job! You guess it!")
			os.Exit(0)
		} else {
			guesses++
			fmt.Printf("You have made %d guesses out of 10.\n", guesses)
		}
	}
	fmt.Printf("Sorry you didn't guess my number. My number was %d.\n", target)
}

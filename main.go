package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const separator = "--------------------"

type JokeResponse struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func fetchJoke() (string, error) {
	resp, err := http.Get("https://official-joke-api.appspot.com/jokes/random")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var joke JokeResponse
	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\n%s", joke.Setup, joke.Punchline), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Joke Fetcher")
	fmt.Println(separator)

	for {
		fmt.Println("1. Fetch a joke")
		fmt.Println("2. Exit")
		fmt.Print("Please choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			joke, err := fetchJoke()
			if err != nil {
				fmt.Println("Failed to fetch joke:", err)
				continue
			}
			fmt.Println(separator)
			fmt.Println(joke)
			fmt.Println(separator)
		case "2":
			fmt.Println("Exiting the program. Have a nice day!")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1 or 2.")

		}
	}
}

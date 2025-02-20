package ui

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"joke-fetcher/joke"
)

// separator for better display formatting
const separator = "--------------------"

// fetchJokeAsync launches a goroutine to fetch a joke asynchronously
func fetchJokeAsync(fetcher joke.JokeFetcher, jokeChan *chan string) {
	go func() {
		jokeData, err := fetcher.FetchJoke()
		if err != nil {
			*jokeChan <- fmt.Sprintf("Failed to fetch joke: %v", err)
			return
		}
		*jokeChan <- joke.TransformJoke(jokeData)
	}()
}

// StartCLI launches the command-line interface for the joke fetcher
func StartCLI(fetcher joke.JokeFetcher) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Joke Fetcher CLI")
	fmt.Println(separator)

	for {
		fmt.Println("1. Fetch a joke")
		fmt.Println("2. Exit")
		fmt.Print("Please choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			// Create a buffered channel to prevent blocking
			jokeChan := make(chan string, 1)

			// Use context to set a timeout for fetching the joke
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			fetchJokeAsync(fetcher, &jokeChan) // Start async joke fetch

			fmt.Println("Fetching joke...")

			select {
			case jokeText := <-jokeChan:
				fmt.Println(separator)
				fmt.Println(jokeText)
				fmt.Println(separator)

			case <-ctx.Done():
				fmt.Println(separator)
				fmt.Println("⏳ Timeout: Joke fetching took too long! Please try again.")
				fmt.Println(separator)
			}

			close(jokeChan) // Safe to close as we buffered it

		case "2":
			fmt.Println("Exiting the program. Have a nice day! 🎉")
			return

		default:
			fmt.Println("Invalid choice. Please enter 1 or 2.")
		}
	}
}

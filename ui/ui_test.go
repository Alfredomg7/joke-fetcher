package ui

import (
	"strings"
	"testing"
	"time"

	"joke-fetcher/joke"
)

// Test fetchJokeAsync success case
func TestFetchJokeAsync_Success(t *testing.T) {
	mockFetcher := joke.MockJokeFetcher{ShouldFail: false}
	jokeChan := make(chan string)

	// Call async function
	fetchJokeAsync(mockFetcher, &jokeChan)

	select {
	case jokeText := <-jokeChan:
		if jokeText == "" {
			t.Errorf("Expected joke text, got empty string")
		}
	case <-time.After(1 * time.Second):
		t.Errorf("Timed out waiting for joke text")
	}
}

// Test fetchJokeAsync failure case
func TestFetchJokeAsync_Failure(t *testing.T) {
	mockFetcher := joke.MockJokeFetcher{ShouldFail: true}
	jokeChan := make(chan string)

	// Call async function
	fetchJokeAsync(mockFetcher, &jokeChan)

	select {
	case jokeText := <-jokeChan:
		if jokeText == "" || !strings.Contains(jokeText, "Failed to fetch joke") {
			t.Errorf("Expected error message, got: %s", jokeText)
		}
	case <-time.After(2 * time.Second):
		t.Errorf("Timeout waiting for joke")
	}
}

// Benchmark fetchJokeAsync function
func BenchmarkFetchJokeAsync(b *testing.B) {
	mockFetcher := joke.MockJokeFetcher{ShouldFail: false}

	for i := 0; i < b.N; i++ {
		jokeChan := make(chan string)
		fetchJokeAsync(mockFetcher, &jokeChan)
		<-jokeChan
	}
}

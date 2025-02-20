package joke

import (
	"testing"
)

// Test success case
func TestFetchJoke_Success(t *testing.T) {
	mockFetcher := MockJokeFetcher{ShouldFail: false}
	jokeData, err := mockFetcher.FetchJoke()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if jokeData.Setup == "" || jokeData.Punchline == "" {
		t.Errorf("Expected joke data, got empty setup or punchline")
	}
}

// Test failure case
func TestFetchJoke_Failure(t *testing.T) {
	mockFetcher := MockJokeFetcher{ShouldFail: true}
	_, err := mockFetcher.FetchJoke()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// Benchmark FetchJoke function
func BenchmarkFetchJoke(b *testing.B) {
	mockFetcher := MockJokeFetcher{ShouldFail: false}

	for i := 0; i < b.N; i++ {
		mockFetcher.FetchJoke()
	}
}

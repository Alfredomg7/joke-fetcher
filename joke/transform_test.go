package joke

import (
	"strings"
	"testing"
)

// Test normal joke transformation
func TestTransformJoke_Success(t *testing.T) {
	jokeData := &JokeResponse{
		Setup:     "Why did the chicken cross the road?",
		Punchline: "To get to the other side!",
	}
	result := TransformJoke(jokeData)

	if !strings.Contains(result, "Why did the chicken") {
		t.Errorf("Expected transformed joke, got: %s", result)
	}

	expectedEmoji := false
	for _, emoji := range emojis {
		if strings.Contains(result, emoji) {
			expectedEmoji = true
			break
		}
	}

	if !expectedEmoji {
		t.Errorf("Expected emoji in transformed joke, got: %s", result)
	}
}

// Test nil input handling
func TestTransformJoke_NilInput(t *testing.T) {
	result := TransformJoke(nil)
	if result != "" {
		t.Errorf("Expected empty string, got: %s", result)
	}
}

// Benchmark TransformJoke function
func BenchmarkTransformJoke(b *testing.B) {
	jokeData := &JokeResponse{
		Setup:     "Why did the chicken cross the road?",
		Punchline: "To get to the other side!",
	}

	for i := 0; i < b.N; i++ {
		TransformJoke(jokeData)
	}
}

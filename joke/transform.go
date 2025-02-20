package joke

import (
	"fmt"
	"math/rand"
	"time"
)

// List of funny/laughing-related emojis
var emojis = []string{"😂", "🤣", "😆", "😹", "😁", "🤭", "😜"}

// TransformJoke formats the joke and appends a random emoji at the end
func TransformJoke(jokeData *JokeResponse) string {
	// Handle nil input
	if jokeData == nil {
		return ""
	}

	// Extract joke text
	jokeText := fmt.Sprintf("%s\n%s", jokeData.Setup, jokeData.Punchline)

	// Select a random emoji
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomEmoji := emojis[r.Intn(len(emojis))]

	// Append emoji at the end
	return fmt.Sprintf("%s %s", jokeText, randomEmoji)
}

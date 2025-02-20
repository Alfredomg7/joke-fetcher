package main

import (
	"joke-fetcher/joke"
	"joke-fetcher/ui"
)

func main() {
	// Inject the real API fetcher into the UI
	apiFetcher := joke.APIJokeFetcher{}
	ui.StartCLI(apiFetcher)
}

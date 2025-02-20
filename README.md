## Overview
CLI application build in Go that fetches random jokes from the Official Joke API. The application provides a menu to either fetch a joke or exit the program.

## Features

- **Async Fetching**: Retrieves random jokes from the Official Joke API in a goroutine with a timeout.
- **Emoji Transformation**: Appends a random laughing emoji to each joke.
- **Menu-Driven UI**: Presents options to fetch a joke or exit the application.
- **Error Handling**: Gracefully handles API and network errors.
- **Unit Tests**: Includes comprehensive tests for both joke fetching and UI.
.

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/joke-fetcher.git
    cd joke-fetcher
    ```

2. **Build the application:**

    ```bash
    go build -o joke-fetcher.exe
    ```

3. **Run the application:**

    ```bash
    ./joke-fetcher
    ```

## Usage

When you run the application, you'll be presented with a menu:

1. **Fetch a joke**: Select this option to get a random joke.
2. **Exit**: Select this option to exit the program.

## Testing
To run all tests execute:
```bash
go test ./...
```

## Dependencies
- Go 1.22.4

# Multiplayer Web Game

This project is a multiplayer web game built using Go, SQLite3, chi, templ, HTMX, and Server-Sent Events (SSE). It provides a framework for real-time interactions between players in a web-based environment.

## Project Structure

```
multiplayer-web-game
├── cmd
│   └── server
│       └── main.go          # Entry point of the application
├── internal
│   ├── game
│   │   └── game.go          # Game logic and state management
│   ├── handlers
│   │   └── handlers.go      # HTTP handlers for various routes
│   ├── models
│   │   └── models.go        # Data models for players and game state
│   └── db
│       └── db.go            # SQLite database management
├── templates
│   └── index.templ          # HTML template for the main page
├── static
│   ├── main.js              # Client-side JavaScript
│   └── styles.css           # CSS styles for the web application
├── go.mod                   # Module definition and dependencies
└── README.md                # Project documentation
```

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- SQLite3

### Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   cd multiplayer-web-game
   ```

2. Initialize the Go module:
   ```
   go mod tidy
   ```

3. Set up the SQLite database:
   - Ensure you have SQLite installed and accessible in your environment.

### Running the Application

To start the server, navigate to the `cmd/server` directory and run:

```
go run main.go
```

The server will start on `http://localhost:3000`.

### Usage

- Access the game in your web browser at `http://localhost:3000`.
- Follow the on-screen instructions to join or create a game.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
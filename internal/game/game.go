package game

import (
	"sync"
	"time"
)

// Game represents the state of the multiplayer game.
type Game struct {
	mu       sync.Mutex
	Players  map[string]*Player
	GameState string
	StartTime time.Time
}

// Player represents a player in the game.
type Player struct {
	ID   string
	Name string
	Score int
}

// NewGame initializes a new game instance.
func NewGame() *Game {
	return &Game{
		Players:  make(map[string]*Player),
		GameState: "waiting",
		StartTime: time.Now(),
	}
}

// AddPlayer adds a new player to the game.
func (g *Game) AddPlayer(id, name string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Players[id] = &Player{ID: id, Name: name, Score: 0}
}

// RemovePlayer removes a player from the game.
func (g *Game) RemovePlayer(id string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.Players, id)
}

// StartGame changes the game state to "in progress".
func (g *Game) StartGame() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.GameState = "in progress"
}

// EndGame changes the game state to "finished".
func (g *Game) EndGame() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.GameState = "finished"
}